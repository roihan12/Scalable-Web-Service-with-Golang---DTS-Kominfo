package service

import (
	"time"

	"github.com/roihan12/assignment-2/model"
	"github.com/roihan12/assignment-2/repository"
)

type OrderService interface {
	CreateOrder(order model.OrderCreateRequest) (model.OrderResponse, error)
	GetOrderByID(orderID uint64) (model.OrderResponse, error)
	UpdateOrder(orderID uint64, input model.OrderUpdateRequest) (model.OrderResponse, error)
	DeleteOrder(ID uint64) error
	GetAllOrders() ([]model.OrderResponse, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) *orderService {
	return &orderService{
		orderRepository,
	}
}

func (service *orderService) CreateOrder(input model.OrderCreateRequest) (model.OrderResponse, error) {

	newOrder := model.Order{
		CustomerName: input.CustomerName,
		OrderAt:      time.Now(),
		Items:        input.Items,
	}

	result, err := service.orderRepository.CreateOrder(newOrder)
	if err != nil {
		return model.OrderResponse{}, err
	}

	return model.ToOrderResponse(*result), nil
}

func (service *orderService) GetOrderByID(orderID uint64) (model.OrderResponse, error) {
	response, err := service.orderRepository.FindOrderById(orderID)
	if err != nil {
		return model.OrderResponse{}, err
	}

	return model.ToOrderResponse(*response), nil
}

func (service *orderService) GetAllOrders() ([]model.OrderResponse, error) {
	response, err := service.orderRepository.GetAllOrders()
	if err != nil {
		return []model.OrderResponse{}, err
	}

	orders := []model.OrderResponse{}
	for _, order := range response {
		order := model.ToOrderResponse(order)

		orders = append(orders, order)
	}

	return orders, nil
}

func (service *orderService) DeleteOrder(ID uint64) error {
	_, err := service.orderRepository.FindOrderById(ID)
	if err != nil {
		return err
	}

	err = service.orderRepository.DeleteOrder(ID)
	if err != nil {
		return err
	}

	return nil
}

func (service *orderService) UpdateOrder(ID uint64, input model.OrderUpdateRequest) (model.OrderResponse, error) {
	_, err := service.orderRepository.FindOrderById(ID)
	if err != nil {
		return model.OrderResponse{}, err
	}
	updateOrder := model.Order{
		CustomerName: input.CustomerName,
		OrderAt:      time.Now(),
		Items:        input.Items,
	}

	updateResult, err := service.orderRepository.UpdateOrder(ID, updateOrder)
	if err != nil {
		return model.OrderResponse{}, err
	}
	return model.ToOrderResponse(*updateResult), nil
}
