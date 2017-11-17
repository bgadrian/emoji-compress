package main

import (
	"net/http"

	"github.com/bgadrian/emoji-compressor/demo/server"
	"google.golang.org/appengine"
)

func main() {
	router := server.NewHandler()
	hl := server.NewLogger(router)

	//TODO find a way to specify the handler and work in APPEngine
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hl.ServeHTTP(w, r)
		// w.Write([]byte("aleluia"))
		return
	})
	appengine.Main()
}
