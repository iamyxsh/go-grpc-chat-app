package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct{}

func (Hash) GenerateHash(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (Hash) CompareHash(plainPassword, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainPassword))
}
