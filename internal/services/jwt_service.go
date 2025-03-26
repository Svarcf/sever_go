package services

import (
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username})
	return token.SignedString([]byte("secret"))
}
