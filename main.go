package main

import (
	"log"
	"net/http"

	"github.com/Pasca11/waterHTTPServer/internal/handler"
	"github.com/Pasca11/waterHTTPServer/internal/repository"
	"github.com/Pasca11/waterHTTPServer/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	// Инициализация слоев
	repo := repository.NewMemoryOrderRepository()
	service := service.NewOrderService(repo)
	handler := handler.NewOrderHandler(service)

	// Настройка маршрутизации
	router := mux.NewRouter()
	router.HandleFunc("/orders", handler.GetOrders).Methods("GET")
	router.HandleFunc("/orders", handler.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", handler.GetOrder).Methods("GET")
	router.HandleFunc("/orders/{id}/status", handler.UpdateOrderStatus).Methods("PUT")

	log.Printf("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
