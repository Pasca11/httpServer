package service

import (
	"strconv"

	"github.com/Pasca11/waterHTTPServer/internal/model"
	"github.com/Pasca11/waterHTTPServer/internal/repository"
)

type OrderService struct {
	repo repository.OrderRepository

	idCounter int
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{
		repo:      repo,
		idCounter: 0,
	}
}

func (s *OrderService) GetOrders() []model.Order {
	return s.repo.GetAll()
}

func (s *OrderService) CreateOrder(order model.Order) model.Order {
	s.idCounter++
	order.ID = strconv.Itoa(s.idCounter)
	return s.repo.Create(order)
}

func (s *OrderService) GetOrder(id string) (model.Order, bool) {
	return s.repo.GetByID(id)
}

func (s *OrderService) UpdateOrderStatus(id string, status string) (model.Order, bool) {
	return s.repo.UpdateStatus(id, status)
}
