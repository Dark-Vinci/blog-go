package helper

import (
	"os"
)

func GetSecretKey() string {
	jwtKey := os.Getenv("JWT_SECRET_KEY")

	if jwtKey == "" {
		jwtKey = "12345678"
	}

	return jwtKey
}