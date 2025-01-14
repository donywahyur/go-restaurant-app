package rest

import (
	"encoding/json"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *userHandler {
	return &userHandler{usecase}
}

func (h *userHandler) RegisterUser(c echo.Context) error {
	var request model.RegisterRequest

	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userCreated, err := h.usecase.RegisterUser(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": userCreated,
	})
}

func (h *userHandler) LoginUser(c echo.Context) error {
	var request model.LoginRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	createdSession, err := h.usecase.LoginUser(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": createdSession,
	})
}
