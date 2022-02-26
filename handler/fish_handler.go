package handler

import (
	"net/http"

	"go_rest_api/model"
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

// @Summary      魚の追加
// @Description  魚を追加し、追加後の魚を返す
// @tags         fishes
// @Accept       json
// @Produce      json
// @Success      200 {object} model.Fish
// @failure      500 {string} string
// @Router       /fishes [post]
// @Param        fish_parameter body model.PostFish true "全ての必須項目"
func FishCreate(c echo.Context) error {
	var fish *model.Fish

	if err := c.Bind(&fish); err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := repository.FishCreate(fish)
	if err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	id, _ := result.LastInsertId()

	fish.ID = int(id)

	return c.JSON(http.StatusOK, fish)
}
