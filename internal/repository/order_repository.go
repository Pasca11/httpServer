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
	orders := []model.Order{
		{
			ID:              "1",
			CustomerName:    "John Doe",
			DeliveryAddress: "123 Main St",
			BottlesCount:    10,
			PhoneNumber:     "1234567890",
		},
		{
			ID:              "2",
			CustomerName:    "Jane Smith",
			DeliveryAddress: "456 Oak Ave",
			BottlesCount:    5,
			PhoneNumber:     "0987654321",
		},
		{
			ID:              "3",
			CustomerName:    "Alice Johnson",
			DeliveryAddress: "789 Pine Rd",
			BottlesCount:    20,
			PhoneNumber:     "1122334455",
		},
		{
			ID:              "4",
			CustomerName:    "Bob Brown",
			DeliveryAddress: "101 Maple St",
			BottlesCount:    15,
			PhoneNumber:     "9988776655",
		},
		{
			ID:              "5",
			CustomerName:    "Charlie Davis",
			DeliveryAddress: "222 Cedar Ave",
			BottlesCount:    30,
			PhoneNumber:     "5554443322",
		},
	}

	return orders
}

func (r *MemoryOrderRepository) GetByID(id string) (model.Order, bool) {
	order, ok := r.orders[id]
	return order, ok
}

func (r *MemoryOrderRepository) Create(order model.Order) model.Order {
	return order
}
