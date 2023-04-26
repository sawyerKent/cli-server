package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Heartbeat(c echo.Context) error {
	response := TextResponse{
		Message: "still breathing",
	}

	return c.JSON(http.StatusOK, response)
}
