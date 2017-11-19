package server

import (
	"github.com/bgadrian/emoji-compress/lz78"

	"github.com/bgadrian/emoji-compress/dictionary"

	"github.com/bgadrian/emoji-compress/bytesmap"
)

func handleBytesmap(op int, req *Request, resp *Response) (err error) {
	switch op {
	case OperationEncode:
		resp.Result, err = bytesmap.EncodeString(req.Text)
	case OperationDecode:
		resp.Result, err = bytesmap.DecodeString(req.Text)
	default:
		err = errMethod
	}
	return
}

func handleDictionary(op int, req *Request, resp *Response) (err error) {
	switch op {
	case OperationEncode:
		// var compressed *dictionary.Result
		resp.Result, err = dictionary.CompressString(req.Text)
		// resp.Result = compressed
	case OperationDecode:
		resp.Result, err = dictionary.DecompressString(req.Dict, req.Text)
	default:
		err = errMethod
	}
	return
}

func handleLZ78(op int, req *Request, resp *Response) (err error) {
	switch op {
	case OperationEncode:
		resp.Result, err = lz78.CompressString(req.Text)
	case OperationDecode:
		resp.Result, err = lz78.DecompressString(req.Text)
	default:
		err = errMethod
	}
	return
}
