package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Pasca11/waterHTTPServer/internal/handler"
	"github.com/Pasca11/waterHTTPServer/internal/repository"
	"github.com/Pasca11/waterHTTPServer/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewMemoryOrderRepository()
	service := service.NewOrderService(repo)
	handler := handler.NewOrderHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/orders", handler.GetOrders).Methods("GET")
	router.HandleFunc("/orders", handler.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", handler.GetOrder).Methods("GET")
	router.HandleFunc("/orders/{id}/status", handler.UpdateOrderStatus).Methods("PUT")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Printf("Сервер запущен на порту 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Получен сигнал завершения, начинаю graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при graceful shutdown: %v", err)
	}

	log.Println("Сервер успешно остановлен")
}
