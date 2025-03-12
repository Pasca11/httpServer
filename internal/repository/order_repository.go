package repository

import (
	"github.com/Pasca11/waterHTTPServer/internal/model"
)

type OrderRepository interface {
	GetAll() []model.Order
	GetByID(id string) (model.Order, bool)
	Create(order model.Order) model.Order
}

type MemoryOrderRepository struct {
	orders map[string]model.Order
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		orders: make(map[string]model.Order),
	}
}

func (r *MemoryOrderRepository) GetAll() []model.Order {
	orders := make([]model.Order, 0, len(r.orders))
	for _, order := range r.orders {
		orders = append(orders, order)
	}
	return orders
}

func (r *MemoryOrderRepository) GetByID(id string) (model.Order, bool) {
	order, ok := r.orders[id]
	return order, ok
}

func (r *MemoryOrderRepository) Create(order model.Order) model.Order {
	r.orders[order.ID] = order
	return order
}
