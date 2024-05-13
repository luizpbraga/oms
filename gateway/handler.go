package main

import "net/http"

type handler struct {
	// todo: gateway dependencies
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRouters(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/custumers/{customerID}/orders", h.handleCreateOrder)
}

func (h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
}
