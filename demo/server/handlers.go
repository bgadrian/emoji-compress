package server

import (
	"net/http"

	"github.com/bgadrian/emoji-compressor/lz78"

	"github.com/bgadrian/emoji-compressor/dictionary"

	"github.com/bgadrian/emoji-compressor/bytesmap"
)

func handleBytesmap(req *Request, resp *Response, r *http.Request) (err error) {
	switch r.Method {
	case http.MethodGet:
		resp.Result, err = bytesmap.EncodeString(req.Text)
	case http.MethodPost:
		resp.Result, err = bytesmap.DecodeString(req.Text)
	default:
		err = errMethod
	}
	return
}

func handleDictionary(req *Request, resp *Response, r *http.Request) (err error) {
	switch r.Method {
	case http.MethodGet:
		// var compressed *dictionary.Result
		resp.Result, err = dictionary.CompressString(req.Text)
		// resp.Result = compressed
	case http.MethodPost:
		resp.Result, err = dictionary.DecompressString(req.Dict, req.Text)
	default:
		err = errMethod
	}
	return
}

func handleLZ78(req *Request, resp *Response, r *http.Request) (err error) {
	switch r.Method {
	case http.MethodGet:
		resp.Result, err = lz78.CompressString(req.Text)
	// case http.MethodPost:
	// 	resp.Result, err = lz78.Decompress(req.Text)
	default:
		err = errMethod
	}
	return
}
