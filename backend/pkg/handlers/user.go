package handlers

import (
	"encoding/json"
	"github.com/aadi-1024/StraySafe/backend/pkg/database"
	"github.com/aadi-1024/StraySafe/backend/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
	"time"
)

func IncidentPostHandler(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("typ").(string) != "usr" && c.Get("typ") != "adm" {
			return c.JSON(http.StatusUnauthorized, models.JsonResponse{
				Message: "user expected",
				Content: nil,
			})
		}

		latitude := c.FormValue("latitude")
		lat, _ := strconv.ParseFloat(latitude, 32)

		longitude := c.FormValue("longitude")
		lon, _ := strconv.ParseFloat(longitude, 32)

		image, _ := c.FormFile("image")
		buf := make([]byte, image.Size)

		file, err := image.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "error parsing image",
				Content: err.Error(),
			})
		}
		_, err = file.Read(buf)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "error parsing image",
				Content: err.Error(),
			})
		}

		filename := strconv.Itoa(c.Get("id").(int)) + strconv.Itoa(int(time.Now().UnixMilli()))
		f, err := os.Create("img/" + filename)
		if err == nil {
			_, err = f.Write(buf)
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "error creating/writing file",
				Content: err.Error(),
			})
		}

		inc := models.Incident{
			Uid:         c.Get("id").(int),
			Latitude:    float32(lat),
			Longitude:   float32(lon),
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			Image:       filename,
		}

		err = d.NewIncident(inc)
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

func UserIncidents(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("typ").(string) != "usr" && c.Get("typ") != "adm" {
			return c.JSON(http.StatusUnauthorized, models.JsonResponse{
				Message: "user expected",
				Content: nil,
			})
		}
		uid := c.Get("id").(int)
		data, err := d.GetIncidents(uid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: data,
		})
	}
}

type jsonReq struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Number    int     `json:"number"`
}

func NearestNgos(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := &jsonReq{}
		err := json.NewDecoder(c.Request().Body).Decode(payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "couldn't decode body",
				Content: err.Error(),
			})
		}
		m, err := d.GetNearestNgo(payload.Latitude, payload.Longitude, payload.Number)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: m,
		})
	}
}
