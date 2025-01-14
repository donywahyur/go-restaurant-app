package rest

import "github.com/labstack/echo/v4"

func LoadRoutesMenu(e *echo.Echo, handler *menuHandler) {
	e.GET("/menu", handler.GetMenu)
}

func LoadRoutesOrder(e *echo.Echo, handler *orderHandler) {
	e.POST("/order", handler.Order)
	e.GET("/order/:orderID", handler.GetOrderInfo)
}
