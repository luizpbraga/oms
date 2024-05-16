package main

import (
	"context"
	"log"
	"net"

	"github.com/luizpbraga/common"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	grpcAddr := common.Getenv("GRPC_ADDR", "localhost:2000")

	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer listener.Close()

	store := NewStore()
	svc := NewService(store)

	// handle the orders
	initGRPCHandler(grpcServer, svc)
	svc.CreateOrder(context.Background())

	log.Println("GRPC server Started at ", grpcAddr)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}
