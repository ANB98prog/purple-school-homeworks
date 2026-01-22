package service

import (
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/internal/repository"
)

type OrderService interface {
	CreateOrder(order CreateOrder) (CreateOrder, error)
}

type orderService struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
	userRepo    repository.UserRepository
}

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository, userRepo repository.UserRepository) OrderService {
	return &orderService{orderRepo: orderRepo, productRepo: productRepo, userRepo: userRepo}
}

var _ OrderService = (*orderService)(nil)

func (service *orderService) CreateOrder(order CreateOrder) (CreateOrder, error) {
	return CreateOrder{}, nil
}
