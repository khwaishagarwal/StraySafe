package main

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/database"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"time"
)

var app = &App{}

func main() {
	db, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	app.Db = db
	app.JwtSecret = []byte(os.Getenv("JWT_SECRET"))
	e := echo.New()
	SetupRoutes(e)

	srv := http.Server{
		Addr:        "0.0.0.0:8080",
		IdleTimeout: 5 * time.Millisecond,
	}

	if err = e.StartServer(&srv); err != nil {
		log.Fatalln(err)
	}
}
