package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Pasca11/waterHTTPServer/internal/model"
	"github.com/Pasca11/waterHTTPServer/internal/service"
	"github.com/gorilla/mux"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders := h.service.GetOrders()
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order model.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Ошибка при декодировании данных: "+err.Error(), http.StatusBadRequest)
		return
	}
	createdOrder := h.service.CreateOrder(order)
	json.NewEncoder(w).Encode(createdOrder)

}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if order, found := h.service.GetOrder(params["id"]); found {
		json.NewEncoder(w).Encode(order)
		return
	}
	json.NewEncoder(w).Encode(&model.Order{})
}
