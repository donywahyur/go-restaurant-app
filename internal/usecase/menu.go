package usecase

import (
	"context"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/repository"
	tracing "go-restaurant-app/internal/tracing"
)

type MenuUsecase interface {
	GetMenuByType(ctx context.Context, menuType string) ([]model.MenuItem, error)
	GetMenuByOrderCode(ctx context.Context, orderCode string) (model.MenuItem, error)
}

type menuUsecase struct {
	repository repository.MenuRepository
}

func NewMenuUsecase(repository repository.MenuRepository) *menuUsecase {
	return &menuUsecase{repository}
}

func (u *menuUsecase) GetMenuByType(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuByType")
	defer span.End()

	var menu []model.MenuItem

	data, err := u.repository.GetMenuByType(ctx, menuType)
	if err != nil {
		return nil, err
	}

	menu = append(menu, data...)

	return menu, nil
}

func (u *menuUsecase) GetMenuByOrderCode(ctx context.Context, orderCode string) (model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuByOrderCode")
	defer span.End()

	var menu model.MenuItem

	menuOrder, err := u.repository.GetMenuByOrderCode(ctx, orderCode)
	if err != nil {
		return menu, err
	}

	return menuOrder, nil
}
