package repository

import (
	"errors"
	"log"

	"github.com/roihan12/assignment-2/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order model.Order) (*model.Order, error)
	FindOrderById(id uint64) (*model.Order, error)
	GetAllOrders() ([]model.Order, error)
	DeleteOrder(id uint64) error
	UpdateOrder(id uint64, order model.Order) (*model.Order, error)
}

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		DB: db,
	}
}

func (repository *OrderRepositoryImpl) CreateOrder(order model.Order) (*model.Order, error) {
	newOrder := model.Order{
		OrderID:      order.OrderID,
		CustomerName: order.CustomerName,
		OrderAt:      order.OrderAt,
		Items:        order.Items,
	}

	err := repository.DB.Create(&newOrder).Error
	if err != nil {
		log.Fatal("error")
		return nil, err
	}

	return &newOrder, nil
}

func (repository *OrderRepositoryImpl) FindOrderById(id uint64) (*model.Order, error) {
	order := model.Order{}

	err := repository.DB.Debug().Preload("Items").First(&order, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}

	return &order, nil
}

func (repository *OrderRepositoryImpl) GetAllOrders() ([]model.Order, error) {
	orders := []model.Order{}

	err := repository.DB.Preload("Items").Find(&orders).Error
	if err != nil {
		return []model.Order{}, err
	}

	return orders, nil
}

func (repository *OrderRepositoryImpl) DeleteOrder(id uint64) error {
	order := model.Order{}
	err := repository.DB.Where("order_id = ?", id).Delete(&order).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *OrderRepositoryImpl) UpdateOrder(id uint64, updatedOrder model.Order) (*model.Order, error) {
	var order model.Order
	var items []model.Item

	err := repository.DB.Preload("Items").First(&order, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}
	order.CustomerName = updatedOrder.CustomerName
	order.Items = updatedOrder.Items
	items = updatedOrder.Items

	repository.DB.Save(&order)
	repository.DB.Save(&items)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
