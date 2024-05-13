package main

import (
	"log"
	"net/http"
	"os"
)

func addr() string {
	addr := os.Getenv("HTTP_ADDR")

	if addr == "" {
		addr = ":8080"
	}

	return addr
}

// load balancer

func main() {
	httpAddr := addr()
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRouters(mux)

	log.Printf("Starting HTPP server at %s\n", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start the server")
	}
}
