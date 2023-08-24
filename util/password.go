package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreatePasswordHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate password hash: %v", err)
	}
	return string(passwordHash), nil
}

func ValidatePassword(password string, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
