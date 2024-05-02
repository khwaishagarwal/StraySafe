package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Id   int    `json:"id"`
	Type string `json:"typ"`
	jwt.RegisteredClaims
}
