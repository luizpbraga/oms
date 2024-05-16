package main

import (
	"errors"
	"net/http"

	"github.com/luizpbraga/common"
	pb "github.com/luizpbraga/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}

	orderRequest := pb.CreateOrderRequest{
		CustumerID: cuntumerID,
		Item:       items,
	}

	// GRPC ERROR HANDLING
	order, err := h.client.CreateOrder(r.Context(), &orderRequest)
	errStatus := status.Convert(err)
	if errStatus != nil {
		if errStatus.Code() != codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, errStatus.Err())
			return
		}
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	common.WriteJson(w, http.StatusOK, order)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
	if len(items) == 0 {
		return errors.New("Items len must be > 0")
	}

	for _, item := range items {
		if item.ID == "" {
			return errors.New("item ID is required")
		}

		if item.Quantity <= 0 {
			return errors.New("item Quantity is required (> 0)")
		}
	}

	return nil
}
