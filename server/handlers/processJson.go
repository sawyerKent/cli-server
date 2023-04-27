package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func ProcessJson(c echo.Context) error {
	incoming := new(IncomingData)
	if err := json.NewDecoder(c.Request().Body).Decode(incoming); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	returnData := make(ReturnData)
	for group, users := range *incoming {
		returnUsers := make([]ReturnUser, 0, len(users))
		for _, user := range users {
			numericMonth, err := monthToNumeric(user.MonthOfBDate)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid month format")
			}

			returnUser := ReturnUser{
				Name:         user.Name,
				MonthOfBDate:  user.MonthOfBDate,
				NumericMonth: numericMonth,
			}
			returnUsers = append(returnUsers, returnUser)
		}
		returnData[group] = returnUsers
	}

	return c.JSON(http.StatusOK, returnData)
}

func monthToNumeric(month string) (int, error) {
	parsedMonth, err := time.Parse("January", month)
	if err != nil {
		return 0, err
	}
	return int(parsedMonth.Month()), nil
}
