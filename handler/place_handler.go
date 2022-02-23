package handler

import (
	"net/http"

	"go_rest_api/repository"

	"github.com/labstack/echo/v4"
)

// PlaceIndex godoc
// @Summary      釣り場の一覧を返す
// @Description  釣り場の一覧を返す
// @Tags         places
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Place
// @Router       /places [get]
func PlaceIndex(c echo.Context) error {
	places, err := repository.PlaceList()

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, places)
}
