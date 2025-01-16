package rest

import (
	"encoding/json"
	"go-restaurant-app/internal/model"
	tracing "go-restaurant-app/internal/tracing"
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
	ctx, span := tracing.CreateSpan(c.Request().Context(), "RegisterUser")
	defer span.End()

	var request model.RegisterRequest

	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userCreated, err := h.usecase.RegisterUser(ctx, request)
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
	ctx, span := tracing.CreateSpan(c.Request().Context(), "LoginUser")
	defer span.End()

	var request model.LoginRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	createdSession, err := h.usecase.LoginUser(ctx, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": createdSession,
	})
}
