package repository

import (
	"go-restaurant-app/internal/model"

	"gorm.io/gorm"
)

type MenuRepository interface {
	GetMenuByType(menuType string) ([]model.MenuItem, error)
	GetMenuByOrderCode(orderCode string) (model.MenuItem, error)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *menuRepository {
	return &menuRepository{db}
}

func (r *menuRepository) GetMenuByType(menuType string) ([]model.MenuItem, error) {
	var menu []model.MenuItem

	err := r.db.Where("type = ?", model.MenuType(menuType)).Find(&menu).Error

	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (r *menuRepository) GetMenuByOrderCode(orderCode string) (model.MenuItem, error) {
	var menu model.MenuItem

	err := r.db.Where("order_code = ?", orderCode).First(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}
