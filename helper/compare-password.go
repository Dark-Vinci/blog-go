package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashed string, password [] byte) bool {
	byteHashed := [] byte (hashed)

	err := bcrypt.CompareHashAndPassword(byteHashed, password)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}