package model

import (
	"time"
)

type Order struct {
	OrderID      uint64      `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `gorm:"not null" json:"customer_name"`
	OrderAt      time.Time `json:"order_at"`
	Items        []Item    `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"items"`
}

type OrderCreateRequest struct {
	CustomerName string `json:"customer_name"`
	OrderAt      string `json:"order_at"`
	Items        []Item `json:"items"`
}

type OrderUpdateRequest struct {
	CustomerName string `json:"customer_name"`
	OrderAt      string `json:"order_at"`
	Items        []Item `json:"items"`
}

type OrderResponse struct {
	OrderID      uint64   `json:"order_id"`
	CustomerName string `json:"customer_name"`
	OrderAt      string `json:"order_at"`
	Items        []Item `json:"items"`
}

func ToOrderResponse(order Order) OrderResponse {
	orderResponse := OrderResponse{
		OrderID:      order.OrderID,
		CustomerName: order.CustomerName,
		OrderAt:      order.OrderAt.Format(time.RFC3339),
		Items:        order.Items,
	}

	return orderResponse
}
