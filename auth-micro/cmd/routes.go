package main

import (
	"github.com/aadi-1024/auth-micro/pkg/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func PopulateRouter(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.POST("/user/login", handlers.LoginHandler(app.Db, app.Jwt))  //login with an existing account
	e.POST("/user/register", handlers.RegistrationHandler(app.Db)) //register a new user
	e.POST("/user/reset", handlers.ResetPasswordHandler(app.Db))   //reset password
	e.POST("/ngo/login", handlers.NgoLoginHandler(app.Db, app.Jwt))
	e.POST("/ngo/register", handlers.NgoRegistrationHandler(app.Db))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
}
