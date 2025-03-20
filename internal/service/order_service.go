package service

import (
	"strconv"

	pb "github.com/Pasca11/grpcServer/proto/gen"
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

func (s *OrderService) CreateOrder(order *pb.OrderRequest) model.Order {
	s.idCounter++
	newOrder := model.Order{
		ID:              strconv.Itoa(s.idCounter),
		CustomerName:    order.CustomerName,
		DeliveryAddress: order.DeliveryAddress,
		BottlesCount:    order.BottlesCount,
		PhoneNumber:     order.PhoneNumber,
	}
	return s.repo.Create(newOrder)
}

func (s *OrderService) GetOrder(id string) (model.Order, bool) {
	return s.repo.GetByID(id)
}
