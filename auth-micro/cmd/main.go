package main

import (
	"github.com/aadi-1024/auth-micro/pkg/database"
	"github.com/aadi-1024/auth-micro/pkg/jwtUtil"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

var app Config

func main() {
	db, err := database.InitDb()
	if err != nil {
		log.Fatalln(err)
	}
	app.Db = db

	conf := jwtUtil.NewJwtConfig(24 * time.Hour)
	app.Jwt = conf

	e := echo.New()
	PopulateRouter(e)
	if err := e.Start("0.0.0.0:9876"); err != nil {
		log.Fatalln(err)
	}
}
