package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got connection: %s", r.Proto)
		w.Write([]byte("Hello"))
	})
	h1s := &http.Server{
		Addr:    ":4000",
		Handler: h2c.NewHandler(handler, &http2.Server{}),
	}
	log.Printf("Serving on https://0.0.0.0:4000")
	log.Fatal(h1s.ListenAndServe())
}
