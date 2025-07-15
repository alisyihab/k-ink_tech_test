package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJWT(username string)(string, error) {
	_ = godotenv.Load()
	
	claims := jwt.MapClaims{
		"username": "username",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	 secret := []byte(os.Getenv("JWT_SECRET"))

	return token.SignedString(secret)
}

func ValidateJWT(tokenStr string)(*jwt.Token, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}