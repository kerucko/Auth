package jwtauth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kerucko/auth/internal/models"
)

func NewToken(user models.User, app models.App, expiration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.Id
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(expiration).Unix()
	claims["app_id"] = app.Id

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
