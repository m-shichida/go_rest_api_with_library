package handler

import (
	"log"
	"net/http"

	"go_json_api/repository"

	"github.com/labstack/echo/v4"
)

func PlaceIndex(c echo.Context) error {
	places, err := repository.PlaceList()

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, places)
}
