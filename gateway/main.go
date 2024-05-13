package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/luizpbraga/common"
)

func main() {
	httpAddr := common.Getenv("HTTP_ADDR", ":3000")
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRouters(mux)

	log.Printf("Starting HTPP server at %s\n", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start the server")
	}
}
