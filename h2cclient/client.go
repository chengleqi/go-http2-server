package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	client := http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	resp, err := client.Get("http://localhost:4000")
	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}

	fmt.Printf(
		"Got response %d: %s %s\n",
		resp.StatusCode, resp.Proto,
		string(body),
	)
}
