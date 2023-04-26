package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HappyLang(c echo.Context) error {
	response := HappyLangResponse{
		FRVRID:   "1234",
		Language: "ExampleLang",
	}

	return c.JSON(http.StatusOK, response)
}
