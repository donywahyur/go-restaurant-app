package repository

import (
	"context"
	"go-restaurant-app/internal/model"
	tracing "go-restaurant-app/internal/tracing"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrderInfo(ctx context.Context, orderID string) (model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateOrder")
	defer span.End()

	err := r.db.WithContext(ctx).Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *orderRepository) GetOrderInfo(ctx context.Context, orderID string) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateOrder")
	defer span.End()

	var order model.Order

	err := r.db.WithContext(ctx).Where("id = ?", orderID).Preload("ProductOrders").Find(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
