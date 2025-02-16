package repository

import (
	"github.com/Pasca11/waterHTTPServer/internal/model"
)

type OrderRepository interface {
	GetAll() []model.Order
	GetByID(id string) (model.Order, bool)
	Create(order model.Order) model.Order
	UpdateStatus(id string, status string) (model.Order, bool)
}

type MemoryOrderRepository struct {
	orders []model.Order
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		orders: make([]model.Order, 0),
	}
}

func (r *MemoryOrderRepository) GetAll() []model.Order {
	return r.orders
}

func (r *MemoryOrderRepository) GetByID(id string) (model.Order, bool) {
	for _, order := range r.orders {
		if order.ID == id {
			return order, true
		}
	}
	return model.Order{}, false
}

func (r *MemoryOrderRepository) Create(order model.Order) model.Order {
	r.orders = append(r.orders, order)
	return order
}

func (r *MemoryOrderRepository) UpdateStatus(id string, status string) (model.Order, bool) {
	for i, order := range r.orders {
		if order.ID == id {
			r.orders[i].Status = status
			return r.orders[i], true
		}
	}
	return model.Order{}, false
}
