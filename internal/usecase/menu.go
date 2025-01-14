package usecase

import (
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/repository"
)

type MenuUsecase interface {
	GetMenuByType(menuType string) ([]model.MenuItem, error)
	GetMenuByOrderCode(orderCode string) (model.MenuItem, error)
}

type menuUsecase struct {
	repository repository.MenuRepository
}

func NewMenuUsecase(repository repository.MenuRepository) *menuUsecase {
	return &menuUsecase{repository}
}

func (u *menuUsecase) GetMenuByType(menuType string) ([]model.MenuItem, error) {
	var menu []model.MenuItem

	data, err := u.repository.GetMenuByType(menuType)
	if err != nil {
		return nil, err
	}

	menu = append(menu, data...)

	return menu, nil
}

func (u *menuUsecase) GetMenuByOrderCode(orderCode string) (model.MenuItem, error) {
	var menu model.MenuItem

	menuOrder, err := u.repository.GetMenuByOrderCode(orderCode)
	if err != nil {
		return menu, err
	}

	return menuOrder, nil
}
