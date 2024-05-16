package main

import (
	"context"
	"log"

	"github.com/luizpbraga/common"
	pb "github.com/luizpbraga/common/api"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store}
}

func (svc *service) CreateOrder(ctx context.Context) error {
	return svc.store.Create(ctx)
}

func (svc *service) ValidateOrder(ctx context.Context, orderReq *pb.CreateOrderRequest) error {
	if len(orderReq.Item) == 0 {
		return common.ErrNoItems
	}

	items := mergeOrders(orderReq)
	log.Println(items)

	// TODO:; validate the stock server

	return nil
}

func mergeOrders(orderReq *pb.CreateOrderRequest) []*pb.ItemsWithQuantity {
	mergedMap := map[string]*pb.ItemsWithQuantity{}
	for _, item := range orderReq.Item {
		if _, ok := mergedMap[item.ID]; ok {
			mergedMap[item.ID].Quantity += item.Quantity
			continue
		}
		mergedMap[item.ID] = item
	}

	var mergedOrders []*pb.ItemsWithQuantity
	for _, value := range mergedMap {
		mergedOrders = append(mergedOrders, value)
	}

	return mergedOrders
}
