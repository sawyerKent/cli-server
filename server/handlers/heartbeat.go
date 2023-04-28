package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawyerKent/cli-server/server/models"
)

func Heartbeat(c echo.Context) error {
	response := models.TextResponse{
		Message: "still breathing",
	}

	return c.JSON(http.StatusOK, response)
}
