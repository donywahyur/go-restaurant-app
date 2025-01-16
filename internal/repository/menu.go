package repository

import (
	"context"
	"go-restaurant-app/internal/model"
	tracing "go-restaurant-app/internal/tracing"

	"gorm.io/gorm"
)

type MenuRepository interface {
	GetMenuByType(ctx context.Context, menuType string) ([]model.MenuItem, error)
	GetMenuByOrderCode(ctx context.Context, orderCode string) (model.MenuItem, error)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *menuRepository {
	return &menuRepository{db}
}

func (r *menuRepository) GetMenuByType(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuByType")
	defer span.End()
	var menu []model.MenuItem

	err := r.db.WithContext(ctx).Where("type = ?", model.MenuType(menuType)).Find(&menu).Error

	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (r *menuRepository) GetMenuByOrderCode(ctx context.Context, orderCode string) (model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuByOrderCode")
	defer span.End()

	var menu model.MenuItem

	err := r.db.WithContext(ctx).Where("order_code = ?", orderCode).First(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}
