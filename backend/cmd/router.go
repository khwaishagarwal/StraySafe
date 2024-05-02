package main

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func SetupRoutes(e *echo.Echo) {
	e.Use(middleware.CORS())
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})
	e.POST("/user/incident", handlers.IncidentPostHandler(app.Db), JwtMiddleware) //new incident
	e.GET("/user/incidents", handlers.UserIncidents(app.Db), JwtMiddleware)       // get all incidents reported by user
	e.GET("/ngo/resolveIncident/:id", handlers.MarkResolvedHandler(app.Db), JwtMiddleware)
	e.GET("/ngo/about/:id", handlers.GetNgoByIdHandler(app.Db))
	e.POST("/ngo/nearest", handlers.NearestNgos(app.Db))
	e.GET("/ngo/dashboard", handlers.GetDashboardHandler(app.Db), JwtMiddleware)
	e.GET("/ngo/incident/:id", handlers.GetIncidentHandler(app.Db), JwtMiddleware)
	//e.GET("/ngo/incident/img/:filename", handlers.GetImage(), JwtMiddleware)
	e.GET("/admin/delete/:id", handlers.DeleteUser(app.Db), JwtMiddleware)
	e.GET("/admin/delete_ngo/:id", handlers.DeleteNgo(app.Db), JwtMiddleware)
	e.GET("/admin/make/:id", handlers.MakeAdmin(app.Db), JwtMiddleware) //make admin
	e.GET("/admin/resolve/:id", handlers.MarkResolvedHandler(app.Db), JwtMiddleware)

	e.Static("static/", "img/")
}
