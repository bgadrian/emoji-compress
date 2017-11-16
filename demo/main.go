package main

import (
	"log"
	"net/http"

	"github.com/bgadrian/emoji-compressor/demo/server"
)

func main() {
	router := server.NewHandler()
	hl := server.NewLogger(router)
	err := http.ListenAndServe(":80", hl)
	log.Fatal(err)
}
