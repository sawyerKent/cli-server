package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HiThere(c echo.Context) error {
	response := TextResponse{
		Message: "Hi there",
	}

	return c.JSON(http.StatusOK, response)
}
