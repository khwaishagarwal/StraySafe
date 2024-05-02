package handlers

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/database"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func DeleteUser(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("typ") != "adm" {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err = db.DeleteUser(id); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "successful")
	}
}

func MakeAdmin(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("typ") != "adm" {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err = db.MakeAdmin(id); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "successful")
	}
}

func DeleteNgo(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("typ") != "adm" {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err = db.DeleteNgo(id); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "successful")
	}
}
