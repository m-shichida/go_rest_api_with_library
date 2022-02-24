package handler

import (
	"net/http"

	"go_rest_api/repository"

	"github.com/labstack/echo/v4"
)

// FishIndex godoc
// @Summary      魚の一覧を返す
// @Description  魚の一覧を返す
// @Tags         fishes
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Fish
// @Router       /fishes [get]
func FishIndex(c echo.Context) error {
	fishes, err := repository.FishList()

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, fishes)
}
