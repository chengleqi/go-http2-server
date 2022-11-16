package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":4000", Handler: http.HandlerFunc(handle)}
	log.Printf("Serving on https://0.0.0.0:4000")
	log.Fatal(srv.ListenAndServeTLS("../cert/cert.pem", "../cert/key.pem"))
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)
	w.Write([]byte("Hello"))
}

// // server push
// func handle(w http.ResponseWriter, r *http.Request) {
// 	// Log the request protocol
// 	log.Printf("Got connection: %s", r.Proto)
// 	// Handle 2nd request, must be before push to prevent recursive calls.
// 	// Don't worry - Go protect us from recursive push by panicking.
// 	if r.URL.Path == "/2nd" {
// 		log.Println("Handling 2nd")
// 		w.Write([]byte("Hello Again!"))
// 		return
// 	}

// 	// Handle 1st request
// 	log.Println("Handling 1st")
// 	// Server push must be before response body is being written.
// 	// In order to check if the connection supports push, we should use
// 	// a type-assertion on the response writer.
// 	// If the connection does not support server push, or that the push
// 	// fails we just ignore it - server pushes are only here to improve
// 	// the performance for HTTP/2 clients.
// 	pusher, ok := w.(http.Pusher)

// 	if !ok {
// 		log.Println("Can't push to client")
// 	} else {
// 		err := pusher.Push("/2nd", nil)
// 		if err != nil {
// 			log.Printf("Failed push: %v", err)
// 		}
// 	}
// 	// Send response body
// 	w.Write([]byte("Hello"))
// }
