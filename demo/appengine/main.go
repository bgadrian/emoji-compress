package main

import (
	"net/http"

	"github.com/bgadrian/emoji-compressor/demo/server"
	"google.golang.org/appengine"
)

func main() {
	router := server.NewHandler()
	hl := server.NewLogger(router)

	//the app engine takes care of the HTTP server
	http.Handle("/", hl)
	appengine.Main()
}
