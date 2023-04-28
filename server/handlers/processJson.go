package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sawyerKent/cli-server/server/models"
)

func ProcessJson(c echo.Context) error {
	incoming := new(models.IncomingData)
	if err := c.Bind(incoming); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	returnData := make(models.ReturnData)
	groups := incoming.GetAllGroups()
	for _, group := range groups {
		userData, _ := incoming.GetUsersInGroup(group)
		for _, user := range userData {
			numericMonth, err := monthToNumeric(user.GetMonthOfBDate())
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid month format")
			}

			returnUser := models.ReturnUser{
				Name:         user.GetName(),
				MonthOfBDate:  user.GetMonthOfBDate(),
				NumericMonth: numericMonth,
			}
			returnData.AddUserToGroup(group, returnUser)
		}
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
