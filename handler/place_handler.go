package handler

import (
	"net/http"

	"go_rest_api/repository"

	"github.com/labstack/echo/v4"
)

func PlaceIndex(c echo.Context) error {
	places, err := repository.PlaceList()

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, places)
}
