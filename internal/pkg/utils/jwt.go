package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/nanfeng/ginchat/internal/config"
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(id string, username string) (string, error) {

	jwt_id, _ := uuid.NewV6()

	cfg := config.Cfg.Jwt

	claims := &CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   id,
			Issuer:    cfg.Iss,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.Exp) * time.Millisecond)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        jwt_id.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.Secret))
}

func ParseToken(tokenString string) (*string, error) {

	secret := config.Cfg.Jwt.Secret

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return &claims.Username, nil
	}

	return nil, errors.New("Invalid token")
}
