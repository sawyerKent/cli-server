package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawyerKent/cli-server/server/models"
)

func HappyLang(c echo.Context) error {
	response := models.HappyLangResponse{
		FRVRID:   "1234",
		Language: "ExampleLang",
	}

	return c.JSON(http.StatusOK, response)
}
