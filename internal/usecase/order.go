package usecase

import (
	"context"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/model/constant"
	"go-restaurant-app/internal/repository"
	tracing "go-restaurant-app/internal/tracing"

	"github.com/google/uuid"
)

type OrderUsecase interface {
	Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error)
	GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error)
}

type orderUsecase struct {
	repository     repository.OrderRepository
	menuRepository repository.MenuRepository
}

func NewOrderUsecase(repository repository.OrderRepository, menuRepository repository.MenuRepository) *orderUsecase {
	return &orderUsecase{repository, menuRepository}
}

func (u *orderUsecase) Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "Order")
	defer span.End()

	order := model.Order{}
	productOrder := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menu, err := u.menuRepository.GetMenuByOrderCode(ctx, orderProduct.OrderCode)
		if err != nil {
			return order, err
		}

		productOrder[i] = model.ProductOrder{
			ID:         uuid.NewString(),
			OrderCode:  menu.OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: menu.Price * int64(menu.Price),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}

	order.ID = uuid.NewString()
	order.ProductOrders = productOrder
	order.Status = constant.OrderStatusProcessed
	order.ReferenceID = request.ReferenceID

	createdOrder, err := u.repository.CreateOrder(ctx, order)
	if err != nil {
		return order, err
	}
	return createdOrder, nil
}

func (u *orderUsecase) GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetOrderInfo")
	defer span.End()

	var order model.Order

	orderFound, err := u.repository.GetOrderInfo(ctx, request.OrderID)
	if err != nil {
		return order, nil
	}

	return orderFound, nil
}
