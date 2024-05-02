package handlers

import (
	"encoding/json"
	"github.com/aadi-1024/auth-micro/pkg/database"
	"github.com/aadi-1024/auth-micro/pkg/jwtUtil"
	"github.com/aadi-1024/auth-micro/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegistrationHandler(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{}
		err := json.NewDecoder(c.Request().Body).Decode(&user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "couldn't parse request body",
				Content: err.Error(),
			})
		}
		err = d.RegisterUser(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: nil,
		})
	}
}

func LoginHandler(d *database.Database, j *jwtUtil.JwtConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{}
		err := json.NewDecoder(c.Request().Body).Decode(&user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "couldn't parse request body",
				Content: err.Error(),
			})
		}
		token, err := d.VerifyLogin(user, j)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.JsonResponse{
				Message: "couldn't verify login information",
				Content: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: token,
		})
	}
}

//type ngoLoginReq struct {
//	Id       int    `json:"id"`
//	Password string `json:"password"`
//}

func NgoLoginHandler(d *database.Database, j *jwtUtil.JwtConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := models.Ngo{}
		err := json.NewDecoder(c.Request().Body).Decode(&payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "couldn't parse request body",
				Content: err.Error(),
			})
		}
		token, err := d.VerifyNgoLogin(payload, j)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: token,
		})
	}
}

func NgoRegistrationHandler(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		ngo := models.Ngo{}
		err := json.NewDecoder(c.Request().Body).Decode(&ngo)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "couldn't parse request body",
				Content: err.Error(),
			})
		}
		id, err := d.RegisterNgo(ngo)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: id,
		})
	}
}

type resetPasswordReq struct {
	models.User
	NewPass string `json:"newPass"`
}

func ResetPasswordHandler(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := resetPasswordReq{}
		err := json.NewDecoder(c.Request().Body).Decode(&payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "couldn't parse request body",
				Content: err.Error(),
			})
		}

		err = d.ResetPassword(models.User{
			Username: payload.Username,
			Password: payload.Password,
		}, []byte(payload.NewPass))
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: nil,
		})
	}
}
