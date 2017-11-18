package server

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
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

var maxPayloadBytes int64 = 10000
var errMethod = errors.New("Request GET for compress, POST for decompress, payload: {text:\"youtetext\"}")
var err404 = errors.New("Supported resources: /bytesmap, /dictionary, /lz78")

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
	resp.Ok = true

	values := r.URL.Query()
	payload := values.Get("payload")
	//anti-flood
	reader := io.LimitReader(strings.NewReader(payload), maxPayloadBytes)
	baseDecoder := base64.NewDecoder(base64.URLEncoding, reader)
	jsonDecoder := json.NewDecoder(baseDecoder)
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
	case "/bytesmap":
		err = handleBytesmap(&req, &resp, r)
	case "/dictionary":
		err = handleDictionary(&req, &resp, r)
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
