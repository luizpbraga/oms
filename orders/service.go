package main

import "context"

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store}
}

func (svc *service) CreateOrder(ctx context.Context) error {
	return svc.store.Create(ctx)
}
