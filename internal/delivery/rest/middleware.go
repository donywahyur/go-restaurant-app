package rest

import (
	"context"
	"errors"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/model/constant"
	"go-restaurant-app/internal/usecase"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadMiddleware(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
}

type authMiddleware struct {
	userUsecase usecase.UserUsecase
}

func NewAuthMiddleware(userUsecase usecase.UserUsecase) *authMiddleware {
	return &authMiddleware{userUsecase}
}
func (am *authMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionData, err := getSessionData(c.Request())
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}
		}
		userID, err := am.userUsecase.CheckSession(sessionData)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}
		}

		authContext := context.WithValue(c.Request().Context(), constant.AuthContextKey, userID)
		c.SetRequest(c.Request().WithContext(authContext))

		return next(c)
	}
}

func getSessionData(r *http.Request) (model.UserSession, error) {
	var userSession model.UserSession

	authString := r.Header.Get("Authorization")
	splitString := strings.Split(authString, " ")

	if len(splitString) != 2 {
		return userSession, errors.New("token not valid")
	}

	accessString := splitString[1]

	userSession.JWTToken = accessString

	return userSession, nil
}
