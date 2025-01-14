package main

import (
	"go-restaurant-app/internal/database"
	"go-restaurant-app/internal/delivery/rest"
	"go-restaurant-app/internal/repository"
	"go-restaurant-app/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := database.GetDB()

	menuRepository := repository.NewMenuRepository(db)
	menuUsecase := usecase.NewMenuUsecase(menuRepository)
	menuHandler := rest.NewMenuHandler(menuUsecase)

	orderRepository := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(orderRepository, menuRepository)
	orderHandler := rest.NewOrderHandler(orderUsecase)

	userRepository := repository.NewUserRepository(db, 64*1024, 4, 32, 12)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := rest.NewUserHandler(userUsecase)

	rest.LoadMiddleware(e)
	rest.LoadRoutesMenu(e, menuHandler)
	rest.LoadRoutesOrder(e, orderHandler)
	rest.LoadRoutesUser(e, userHandler)

	e.Logger.Fatal(e.Start(":8080"))

}
