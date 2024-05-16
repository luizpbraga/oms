package main

import (
	"context"
	"log"

	pb "github.com/luizpbraga/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrdersService
}

func initGRPCHandler(grpcServer *grpc.Server, service OrdersService) {
	handler := &grpcHandler{service: service}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, orderReq *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New Order Received %v\n", orderReq)

	h.service.ValidateOrder(ctx, orderReq)

	order := pb.Order{ID: "32"}

	return &order, nil
}
