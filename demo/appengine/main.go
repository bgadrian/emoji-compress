package main

import (
	"log"
	"net/http"

	"github.com/bgadrian/emoji-compressor/demo/server"
	"google.golang.org/appengine"
)

func main() {
	router := server.NewHandler()
	hl := server.NewLogger(router)
	go func() {
		err := http.ListenAndServe(":80", hl)
		log.Fatal(err)
	}()
	appengine.Main()
}
