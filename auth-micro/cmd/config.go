package main

import (
	"github.com/aadi-1024/auth-micro/pkg/database"
	"github.com/aadi-1024/auth-micro/pkg/jwtUtil"
	"github.com/aadi-1024/auth-micro/pkg/mail"
)

type Config struct {
	Db   *database.Database
	Jwt  *jwtUtil.JwtConfig
	Mail *mail.Mail
}
