package handlers

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/database"
	"github.com/aadi-1024/StraySafe/backend/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetNgoByIdHandler(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "invalid id",
				Content: err.Error(),
			})
		}
		ngo, err := d.GetNgoById(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}

		num, err := d.GetResolvedCases(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}

		m := make(map[string]any)
		m["ngo"] = ngo
		m["num"] = num

		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: m,
		})

	}
}
