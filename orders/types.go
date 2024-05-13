package main

import "context"

type OrdersService interface {
	// grpc + payload
	Create(context.Context) error
}

type OrdersStore interface {
	Create(context.Context) error
}
