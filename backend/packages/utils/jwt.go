package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct{}

type Claims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func (JWT) GenerateToken(id, secret string, exp time.Time) (string, error) {
	claims := &Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(token, secret string) (string, bool) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}
		return []byte(secret), nil
	})

	var id string

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		id = claims["id"].(string)
	} else {
		fmt.Println(err)
	}

	if err != nil {
		return "", false
	}

	return id, true
}
