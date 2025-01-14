package usecase

import (
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/model/constant"
	"go-restaurant-app/internal/repository"

	"github.com/google/uuid"
)

type OrderUsecase interface {
	Order(request model.OrderMenuRequest) (model.Order, error)
	GetOrderInfo(request model.GetOrderInfoRequest) (model.Order, error)
}

type orderUsecase struct {
	repository     repository.OrderRepository
	menuRepository repository.MenuRepository
}

func NewOrderUsecase(repository repository.OrderRepository, menuRepository repository.MenuRepository) *orderUsecase {
	return &orderUsecase{repository, menuRepository}
}

func (u *orderUsecase) Order(request model.OrderMenuRequest) (model.Order, error) {

	order := model.Order{}
	productOrder := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menu, err := u.menuRepository.GetMenuByOrderCode(orderProduct.OrderCode)
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

	createdOrder, err := u.repository.CreateOrder(order)
	if err != nil {
		return order, err
	}
	return createdOrder, nil
}

func (u *orderUsecase) GetOrderInfo(request model.GetOrderInfoRequest) (model.Order, error) {
	var order model.Order

	orderFound, err := u.repository.GetOrderInfo(request.OrderID)
	if err != nil {
		return order, nil
	}

	return orderFound, nil
}
