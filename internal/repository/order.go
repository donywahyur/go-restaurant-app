package repository

import (
	"go-restaurant-app/internal/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrderInfo(orderID string) (model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(order model.Order) (model.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *orderRepository) GetOrderInfo(orderID string) (model.Order, error) {
	var order model.Order

	err := r.db.Where("id = ?", orderID).Preload("ProductOrders").Find(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
