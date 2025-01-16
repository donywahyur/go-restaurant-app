package usecase

import (
	"context"
	"errors"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/repository"
	tracing "go-restaurant-app/internal/tracing"

	"github.com/google/uuid"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error)
	LoginUser(ctx context.Context, request model.LoginRequest) (model.UserSession, error)
	CheckSession(ctx context.Context, userSession model.UserSession) (string, error)
}

type userUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repository repository.UserRepository) *userUsecase {
	return &userUsecase{
		repository,
	}
}

func (u *userUsecase) RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	var user model.User
	username := request.Username
	password := request.Password

	registeredUser, err := u.repository.CheckRegistered(ctx, username)
	if err != nil {
		return user, err
	}

	if registeredUser {
		return user, errors.New("user already registered")
	}

	passwordHash, err := u.repository.GenerateUserHash(ctx, password)
	if err != nil {
		return user, err
	}

	user.ID = uuid.NewString()
	user.Username = username
	user.HashPassword = passwordHash

	userCreated, err := u.repository.RegisterUser(ctx, user)
	if err != nil {
		return user, err
	}

	return userCreated, nil
}

func (u *userUsecase) LoginUser(ctx context.Context, request model.LoginRequest) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "LoginUser")
	defer span.End()

	var userSession model.UserSession

	username := request.Username
	password := request.Password

	userData, err := u.repository.GetUserData(ctx, username)
	if err != nil {
		return userSession, err
	}

	if userData.Username == "" {
		return userSession, errors.New("user not found")
	}

	compareHash, err := u.repository.CompareHash(ctx, password, userData.HashPassword)
	if err != nil {
		return userSession, err
	}

	if !compareHash {
		return userSession, errors.New("password doesnot match")
	}

	createdSession, err := u.repository.CreateUserSession(ctx, userData.ID)
	if err != nil {
		return userSession, err
	}

	return createdSession, nil
}

func (u *userUsecase) CheckSession(ctx context.Context, userSession model.UserSession) (string, error) {
	ctx, span := tracing.CreateSpan(ctx, "LoginUser")
	defer span.End()

	userID, err := u.repository.CheckSession(ctx, userSession)
	if err != nil {
		return "", err
	}

	userData, err := u.repository.GetUserDataByID(ctx, userID)
	if err != nil {
		return "", err
	}

	if userData.ID == "" {
		return "", errors.New("user not found")
	}

	return userID, nil
}
