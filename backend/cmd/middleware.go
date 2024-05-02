package main

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearerToken := c.Request().Header.Get("Authorization")
		if bearerToken == "" {
			return c.JSON(http.StatusUnauthorized, models.JsonResponse{
				Message: "expected jwt token",
				Content: nil,
			})
		}

		var token string
		spl := strings.Split(bearerToken, " ")
		if len(spl) == 1 {
			token = spl[0]
		} else {
			token = spl[1]
		}

		clms := models.Claims{}
		tkn, err := jwt.ParseWithClaims(token, &clms, func(token *jwt.Token) (interface{}, error) {
			return app.JwtSecret, nil
		})
		if !tkn.Valid {
			return c.JSON(http.StatusUnauthorized, models.JsonResponse{
				Message: "invalid token",
				Content: nil,
			})
		}
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.JsonResponse{
				Message: "something went wrong",
				Content: err.Error(),
			})
		}
		if clms.ExpiresAt.Before(time.Now()) {
			return c.JSON(http.StatusUnauthorized, models.JsonResponse{
				Message: "session ended",
				Content: nil,
			})
		}

		c.Set("typ", clms.Type)
		c.Set("id", clms.Id)
		return next(c)
	}
}
