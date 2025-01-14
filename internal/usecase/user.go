package usecase

import (
	"errors"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/repository"

	"github.com/google/uuid"
)

type UserUsecase interface {
	RegisterUser(request model.RegisterRequest) (model.User, error)
	LoginUser(request model.LoginRequest) (model.UserSession, error)
}

type userUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repository repository.UserRepository) *userUsecase {
	return &userUsecase{
		repository,
	}
}

func (u *userUsecase) RegisterUser(request model.RegisterRequest) (model.User, error) {
	var user model.User
	username := request.Username
	password := request.Password

	registeredUser, err := u.repository.CheckRegistered(username)
	if err != nil {
		return user, err
	}

	if registeredUser {
		return user, errors.New("user already registered")
	}

	passwordHash, err := u.repository.GenerateUserHash(password)
	if err != nil {
		return user, err
	}

	user.ID = uuid.NewString()
	user.Username = username
	user.HashPassword = passwordHash

	userCreated, err := u.repository.RegisterUser(user)
	if err != nil {
		return user, err
	}

	return userCreated, nil
}

func (u *userUsecase) LoginUser(request model.LoginRequest) (model.UserSession, error) {
	var userSession model.UserSession

	username := request.Username
	password := request.Password

	userData, err := u.repository.GetUserData(username)
	if err != nil {
		return userSession, err
	}

	if userData.Username == "" {
		return userSession, errors.New("user not found")
	}

	compareHash, err := u.repository.CompareHash(password, userData.HashPassword)
	if err != nil {
		return userSession, err
	}

	if !compareHash {
		return userSession, errors.New("password doesnot match")
	}

	createdSession, err := u.repository.CreateUserSession(userData.ID)
	if err != nil {
		return userSession, err
	}

	return createdSession, nil
}
