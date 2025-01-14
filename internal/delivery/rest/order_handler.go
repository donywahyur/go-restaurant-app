package rest

import (
	"encoding/json"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderHandler(orderUsecase usecase.OrderUsecase) *orderHandler {
	return &orderHandler{orderUsecase}
}

func (h *orderHandler) Order(c echo.Context) error {
	var request model.OrderMenuRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	orderData, err := h.orderUsecase.Order(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}

func (h *orderHandler) GetOrderInfo(c echo.Context) error {
	var request model.GetOrderInfoRequest

	orderID := c.Param("orderID")

	request.OrderID = orderID

	orderFound, err := h.orderUsecase.GetOrderInfo(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderFound,
	})
}
