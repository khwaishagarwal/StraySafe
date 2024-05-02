package handlers

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/database"
	"github.com/aadi-1024/StraySafe/backend/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
)

func MarkResolvedHandler(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("typ") != "ngo" && c.Get("typ") != "adm" {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "expected ngo",
				Content: nil,
			})
		}
		//incId, _ := strconv.Atoi(c.FormValue("id"))
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "invalid id",
				Content: err.Error(),
			})
		}
		inc := models.Incident{Id: id}
		err = db.MarkResolved(inc, c.Get("id").(int))

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

// GetDashboardHandler returns the nearest few incidents, along with info about the NGO including
// number of incidents resolved
func GetDashboardHandler(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("typ") != "ngo" && c.Get("typ") != "adm" {
			return c.JSON(http.StatusUnauthorized, models.JsonResponse{
				Message: "ngo expected",
				Content: nil,
			})
		}

		id := c.Get("id").(int)
		ngo, err := db.GetNgoById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}

		num, err := db.GetResolvedCases(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}

		inc, err := db.GetNearestCases(ngo.Latitude, ngo.Longitude, 10)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}

		m := make(map[string]any)
		m["ngo"] = ngo
		m["num"] = num
		m["inc"] = inc

		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: m,
		})
	}
}

func GetIncidentHandler(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("typ") != "ngo" && c.Get("typ") != "adm" {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "ngo expected",
				Content: nil,
			})
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "invalid id",
				Content: err.Error(),
			})
		}

		inc, err := d.GetIncident(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: inc,
		})
	}
}

func GetImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("typ") != "ngo" && c.Get("typ") != "adm" {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "ngo expected",
				Content: nil,
			})
		}

		filename := c.Param("filename")
		_, err := os.Open("img/" + filename)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "image error",
				Content: err.Error(),
			})
		}

		return c.File("img/" + filename)
	}
}
