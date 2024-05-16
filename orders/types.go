package main

import (
	"context"
	pb "github.com/luizpbraga/common/api"
)

type OrdersService interface {
	// grpc + payload
	CreateOrder(context.Context) error
	ValidateOrder(context.Context, *pb.CreateOrderRequest) error
}

type OrdersStore interface {
	Create(context.Context) error
}
