package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/la4zen/balance-platform-hackaton/pkg/models"
)

const (
	DB_DSN = "database=balance user=postgres password=897+897 sslmode=disable"
)

var (
	JWT_KEY = []byte("secretKeys")
)

func GenerateToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		id,
		jwt.StandardClaims{},
	})
	t, err := token.SignedString(JWT_KEY)
	if err != nil {
		return "", err
	}
	return t, nil
}
