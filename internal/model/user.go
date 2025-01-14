package model

import "github.com/golang-jwt/jwt"

type User struct {
	ID           string `json:"id"`
	Username     string `gorm:"unique" json:"username"`
	HashPassword string `json:"-"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSession struct {
	JWTToken string `json:"jwt_token"`
}

type MyClaims struct {
	jwt.StandardClaims
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
