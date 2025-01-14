package constant

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var APPLICATION_NAME = "GO RESTAURANT APP"
var JWT_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret")
