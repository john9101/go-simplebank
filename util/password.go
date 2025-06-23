package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(passowrd string) (string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passowrd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func MatchPassword(passowrd string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passowrd))
}