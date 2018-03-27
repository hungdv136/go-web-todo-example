package common

import (
	"fmt"
	"time"

	"todo/configs"

	jwt "github.com/dgrijalva/jwt-go"
)

func NewWithClaims(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": "admin",
		"sub":  username,
		"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	return token.SignedString([]byte(configs.GetConfig().GetJwtSecretKey()))
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configs.GetConfig().GetJwtSecretKey()), nil
	})
	return err == nil && token.Valid
}
