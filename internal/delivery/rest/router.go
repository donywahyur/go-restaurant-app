package rest

import (
	"github.com/labstack/echo/v4"
)

func LoadRoutesMenu(e *echo.Echo, handler *menuHandler, authMiddleware *authMiddleware) {
	menuGroup := e.Group("/menu")
	menuGroup.GET("", handler.GetMenu, authMiddleware.CheckAuth)
}

func LoadRoutesOrder(e *echo.Echo, handler *orderHandler, authMiddleware *authMiddleware) {
	orderGroup := e.Group("/order")
	orderGroup.POST("", handler.Order, authMiddleware.CheckAuth)
	orderGroup.GET("/:orderID", handler.GetOrderInfo, authMiddleware.CheckAuth)
}

func LoadRoutesUser(e *echo.Echo, handler *userHandler) {
	userGroup := e.Group("/user")
	userGroup.POST("/register", handler.RegisterUser)
	userGroup.POST("/login", handler.LoginUser)
}
