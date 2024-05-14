package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/luizpbraga/common"
	pb "github.com/luizpbraga/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	httpAddr := common.Getenv("HTTP_ADDR", ":8080")
	orderServiceAddr := "localhost:2000"

	conn, err := grpc.Dial(
		orderServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	log.Println("Dialing Order Service At ", orderServiceAddr)

	mux := http.NewServeMux()
	handler := NewHandler(client)
	handler.registerRouters(mux)

	log.Printf("Starting HTTP server at %s\n", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start the server: ", err.Error())
	}
}
