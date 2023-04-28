package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawyerKent/cli-server/server/models"
)

func HappyLangPost(c echo.Context) error {
	contentType := c.Request().Header.Get("Content-Type")

	happyLangRequest := new(models.HappyLangResponse)

	if contentType == "application/x-www-form-urlencoded" {
		happyLangRequest.FRVRID = c.FormValue("FRVRID")
		happyLangRequest.Language = c.FormValue("language")
	} else if contentType == "application/json" {
		if err := json.NewDecoder(c.Request().Body).Decode(happyLangRequest); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Content-Type, only application/x-www-form-urlencoded and application/json are supported")
	}

	response := models.HappyLangResponse{
		FRVRID:   happyLangRequest.FRVRID,
		Language: happyLangRequest.Language,
	}

	return c.JSON(http.StatusOK, response)
}
