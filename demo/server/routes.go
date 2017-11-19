package server

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"google.golang.org/appengine"
)

//Response API base result structure
type Response struct {
	Ok     bool        `json:"ok"`
	Err    string      `json:"err"`
	Result interface{} `json:"response"`
}

//Request API request format
type Request struct {
	Text string            `json:"text"`
	Dict map[string]string `json:"dict"`
}

const (
	OperationEncode = iota //0
	OperationDecode        //1
)

var maxPayloadBytes int64 = 10000
var errMethod = errors.New("We support only POST requests with a json body payload: {text:\"youtetext\"}")
var err404 = errors.New("Supported resources: /bytesmap, /dictionary, /lz78 with /encode and /decode calls.")

type simpleHandler struct {
}

func NewHandler() http.Handler {
	return simpleHandler{}
}

func (h simpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var resp Response
	var req Request

	defer func() {
		if resp.Err == err404.Error() {
			w.WriteHeader(http.StatusNotFound)
		} else if resp.Ok == false {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("error writing response: %v", err)
		}
	}()

	w.Header().Set("Content-Type", "application/json")
	if appengine.IsDevAppServer() {
		w.Header().Add("Access-Control-Allow-Origin", "*")
	} else {
		w.Header().Add("Access-Control-Allow-Origin", "http://emoji-compress.com")
	}

	resp.Ok = true

	//anti-flood
	r.Body = http.MaxBytesReader(w, r.Body, maxPayloadBytes)
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&req)
	if err != nil && err != io.EOF {
		resp.Ok = false
		resp.Err = err.Error()
		log.Println("error decoding request:", err)
		return
	}

	//TODO normalize the UTF-8 to optimize the compressed length?
	// req.Text = norm.NFC.Bytes()

	switch r.URL.Path {
	case "/bytesmap/encode":
		err = handleBytesmap(OperationEncode, &req, &resp)
	case "/bytesmap/decode":
		err = handleBytesmap(OperationDecode, &req, &resp)
	case "/dictionary/encode":
		err = handleDictionary(OperationEncode, &req, &resp)
	case "/dictionary/decode":
		err = handleDictionary(OperationDecode, &req, &resp)
	case "/lz78/encode":
		err = handleLZ78(OperationEncode, &req, &resp)
	case "/lz78/decode":
		err = handleLZ78(OperationDecode, &req, &resp)
	default:
		err = err404
	}

	if err != nil {
		resp.Ok = false
		resp.Err = err.Error()
		log.Println("erro handling: ", err)
		return
	}
	return
}

func NewLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}
