package rest

import (
	"go-restaurant-app/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type menuHandler struct {
	menuUsecase usecase.MenuUsecase
}

func NewMenuHandler(menuUsecase usecase.MenuUsecase) *menuHandler {
	return &menuHandler{menuUsecase}
}

func (h *menuHandler) GetMenu(c echo.Context) error {

	menuType := c.FormValue("menu_type")

	menu, err := h.menuUsecase.GetMenuByType(menuType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menu,
	})
}
