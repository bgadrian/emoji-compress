package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bgadrian/emoji-compressor/dictionary"
)

type one struct {
	method   string
	url      string
	payload  *Request
	status   int
	response *Response
}

var dicres1 = dictionary.Result{
	Ratio:   1.1538461,
	Source:  "AAA is BBB, or BBB is AAA?",
	Archive: "😀 is 😬, or 😬 is 😀?",
	Words: map[string]string{
		"AAA": "😀",
		"BBB": "😬",
	},
}

var table = []one{
	{http.MethodGet, "/xxx", &Request{},
		http.StatusBadRequest, &Response{Ok: false, Err: err404.Error()}},
	{http.MethodPut, "/bytesmap", &Request{Text: "127.0.0.1"},
		http.StatusBadRequest, &Response{Ok: false, Err: errMethod.Error()}},

	{http.MethodGet, "/bytesmap", &Request{Text: "127.0.0.1"},
		http.StatusOK, &Response{Ok: true, Result: "🙇🙈🙍🙀🙆🙀🙆🙀🙇"}},
	{http.MethodPost, "/bytesmap", &Request{Text: "🙇🙈🙍🙀🙆🙀🙆🙀🙇"},
		http.StatusOK, &Response{Ok: true, Result: "127.0.0.1"}},

	{http.MethodGet, "/dictionary", &Request{Text: dicres1.Source},
		http.StatusOK, &Response{Ok: true, Result: dicres1}},
	{http.MethodPost, "/dictionary", &Request{Text: dicres1.Archive, Dict: dicres1.Words},
		http.StatusOK, &Response{Ok: true, Result: dicres1.Source}},
}

func TestTableall(t *testing.T) {
	//thanks to https://elithrar.github.io/article/testing-http-handlers-go/
	for _, test := range table {
		info := test.method + " " + test.url
		resp, err := json.Marshal(test.payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(test.method, test.url, bytes.NewReader(resp))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := NewHandler()
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != test.status {
			t.Fatalf("%s: handler returned wrong status code: got %v want %v",
				info, status, test.status)
		}

		// Check the response body is what we expect.
		expected, err := json.Marshal(test.response)
		if err != nil {
			t.Fatal(err)
		}
		expString := string(expected) + "\n"
		if rr.Body.String() != string(expString) {
			t.Errorf("%s: handler returned unexpected body: got %q want %q",
				info, rr.Body.String(), expString)
		}
	}
}
