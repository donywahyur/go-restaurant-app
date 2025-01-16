package repository

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/model/constant"
	tracing "go-restaurant-app/internal/tracing"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"

	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user model.User) (model.User, error)
	CheckRegistered(ctx context.Context, username string) (bool, error)
	GenerateUserHash(ctx context.Context, password string) (string, error)
	CompareHash(ctx context.Context, password, passwordHash string) (bool, error)
	GetUserData(ctx context.Context, username string) (model.User, error)
	CreateUserSession(ctx context.Context, userID string) (model.UserSession, error)
	CheckSession(ctx context.Context, userSession model.UserSession) (string, error)
}

type userRepository struct {
	db          *gorm.DB
	time        uint32
	memory      uint32
	parallelism uint32
	keyLen      uint32
}

func NewUserRepository(db *gorm.DB,
	time uint32,
	memory uint32,
	parallelism uint32,
	keyLen uint32,
) *userRepository {
	return &userRepository{db, time, memory, parallelism, keyLen}
}

func (r *userRepository) RegisterUser(ctx context.Context, user model.User) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CheckRegistered(ctx context.Context, username string) (bool, error) {

	ctx, span := tracing.CreateSpan(ctx, "CheckRegistered")
	defer span.End()

	var user model.User

	err := r.db.WithContext(ctx).Where("username = ? ", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}

	return user.ID != "", nil
}

func (r *userRepository) GenerateUserHash(ctx context.Context, password string) (string, error) {

	_, span := tracing.CreateSpan(ctx, "GenerateUserHash")
	defer span.End()

	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	argonHash := argon2.IDKey([]byte(password), salt, r.time, r.memory, uint8(r.parallelism), r.keyLen)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(argonHash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, r.memory, r.time, r.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func (r *userRepository) GetUserData(ctx context.Context, username string) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetUserData")
	defer span.End()

	var user model.User

	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CompareHash(ctx context.Context, password, hashPassword string) (bool, error) {

	_, span := tracing.CreateSpan(ctx, "CompareHash")
	defer span.End()

	vals := strings.Split(hashPassword, "$")
	if len(vals) != 6 {
		return false, errors.New("invalid hash")
	}

	var memory, time uint32
	var parallelism uint8

	_, err := fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &memory, &time, &parallelism)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return false, err
	}

	decryptedHash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return false, err
	}

	var keyLen = uint32(len(decryptedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, time, memory, parallelism, keyLen)

	return subtle.ConstantTimeCompare(comparisonHash, decryptedHash) == 1, nil
}

func (r *userRepository) CreateUserSession(ctx context.Context, userID string) (model.UserSession, error) {

	_, span := tracing.CreateSpan(ctx, "CreateUserSession")
	defer span.End()

	var userSession model.UserSession

	accessClaim := model.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    constant.APPLICATION_NAME,
			Subject:   userID,
			ExpiresAt: time.Now().Add(constant.JWT_EXPIRATION_DURATION).Unix(),
		},
	}

	token := jwt.NewWithClaims(
		constant.JWT_SIGNING_METHOD,
		accessClaim,
	)

	signedToken, err := token.SignedString(constant.JWT_SIGNATURE_KEY)
	if err != nil {
		return userSession, err
	}

	userSession.JWTToken = signedToken

	return userSession, nil
}

func (r *userRepository) CheckSession(ctx context.Context, userSession model.UserSession) (string, error) {

	_, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()

	accessToken, err := jwt.ParseWithClaims(userSession.JWTToken, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != constant.JWT_SIGNING_METHOD {
			return nil, errors.New("signing method invalid")
		}

		return constant.JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return "", err
	}

	accessTokenClaim, ok := accessToken.Claims.(*model.MyClaims)
	if !ok || !accessToken.Valid {
		return "", errors.New("unauthorized")
	}

	return accessTokenClaim.Subject, nil
}
