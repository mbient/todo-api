package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/mbient/todo-api/internal/models"

	"os"
	"time"
)

func CreateJWTToken(email string) (string, error) {

	claims := &models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "todo-api",
			Subject:   email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return tokenString, err
}
