package service

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTService struct {
	secretKey string
	publicKey string
}

type JWTClaims struct {
	sub string
	exp string
	iat string
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{secretKey: secret}
}

func (s JWTService) Create(id int) (string, error) {
	currTime := time.Now

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": currTime().Add(3 * time.Hour).Unix(),
		"iat": currTime().Unix(),
	})

	return token.SignedString([]byte(s.secretKey))
}
