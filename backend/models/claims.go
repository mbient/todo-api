package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims // contains standard claims like exp, iat, iss, etc.
}
