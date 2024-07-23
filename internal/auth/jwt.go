package auth

import (
	"fmt"
	"go-api-starter/internal/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "pass"

type JWT struct{}

func NewJWT() JWT {
	return JWT{}
}

func (j JWT) Authenticate(user types.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       user.ID,
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JWT) Authorize(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
