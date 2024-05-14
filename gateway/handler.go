package main

import (
	"log"
	"net/http"

	"github.com/luizpbraga/common"
	pb "github.com/luizpbraga/common/api"
)

type handler struct {
	// todo: gateway dependencies
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{
		client: client,
	}
}

func (h *handler) registerRouters(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/custumers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	cuntumerID := r.PathValue("curtomerID")

	var items []*pb.ItemsWithQuantity

	if err := common.ReadFromJson(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	orderRequest := pb.CreateOrderRequest{
		CustumerID: cuntumerID,
		Item:       items,
	}

	h.client.CreateOrder(r.Context(), &orderRequest)
}
