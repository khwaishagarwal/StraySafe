package jwtUtil

import (
	"github.com/aadi-1024/auth-micro/pkg/models"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type JwtConfig struct {
	jwtSecret []byte
	expiry    time.Duration
}

func NewJwtConfig(exp time.Duration) *JwtConfig {
	return &JwtConfig{
		jwtSecret: []byte(os.Getenv("JWT_SECRET")),
		//jwtSecret: []byte("secret"),
		expiry: exp,
	}
}

func (j *JwtConfig) GenerateToken(uid int, typ string) (string, error) {
	clms := models.Claims{
		uid,
		typ,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiry)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clms)
	signedString, err := token.SignedString(j.jwtSecret)
	if err != nil {
		return "", err
	}
	return signedString, nil
}
