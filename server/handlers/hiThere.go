package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawyerKent/cli-server/server/models"
)

func HiThere(c echo.Context) error {
	response := models.TextResponse{
		Message: "Hi there",
	}

	return c.JSON(http.StatusOK, response)
}
