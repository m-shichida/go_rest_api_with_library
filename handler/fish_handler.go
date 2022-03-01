package handler

import (
	"net/http"
	"strconv"

	"go_rest_api/model"
	"go_rest_api/repository"

	"github.com/labstack/echo/v4"
)

// FishIndex godoc
// @Summary      index
// @Description  魚の情報の一覧を返す
// @Tags         fishes
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Fish
// @Router       /fishes [get]
func FishIndex(c echo.Context) error {
	fishes, err := repository.FishList()

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, fishes)
}

// @Summary      create
// @Description  魚の情報を追加し、追加後の魚を返す
// @tags         fishes
// @Accept       json
// @Produce      json
// @Success      200 {object} model.Fish
// @failure      500 {string} string
// @Router       /fishes [post]
// @Param        fish_parameter body model.FishParameter true "全て必須項目"
func FishCreate(c echo.Context) error {
	var fish *model.Fish

	if err := c.Bind(&fish); err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := repository.FishCreate(fish)
	if err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusInternalServerError, err)
	}

	id, _ := result.LastInsertId()

	fish.Id = int(id)

	return c.JSON(http.StatusOK, fish)
}

// @Summary      show
// @Description  魚の詳細を返す
// @tags         fishes
// @Accept       json
// @Produce      json
// @Success      200 {object} model.Fish
// @Success      404 {string} string
// @failure      500 {string} string
// @Router       /fishes/{id} [get]
// @Param        id path int true "fish ID"
func FishShow(c echo.Context) error {
	var fish *model.Fish

	id, _ := strconv.Atoi(c.Param("id"))
	fish, err := repository.FishGetById(id)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, fish)
}

// @Summary      update
// @Description  魚の情報を更新する
// @tags         fishes
// @Accept       json
// @Produce      json
// @Success      200 {object} model.Fish
// @Success      404 {string} string
// @failure      500 {string} string
// @Router       /fishes/{id} [patch]
// @Param        id path int true "fish ID"
// @Param        fish_parameter body model.FishParameter true "全て必須項目"
func FishUpdate(c echo.Context) error {
	var fish model.Fish

	id, _ := strconv.Atoi(c.Param("id"))
	fish.Id = id

	if err := c.Bind(&fish); err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusBadRequest, err)
	}

	_, err := repository.FishUpdate(&fish)
	if err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, fish)
}

// @Summary      delete
// @Description  魚の情報を削除する
// @tags         fishes
// @Accept       json
// @Produce      json
// @Success      200 {string} string
// @Success      404 {string} string
// @failure      500 {string} string
// @Router       /fishes/{id} [delete]
// @Param        id path int true "fish ID"
func FishDestroy(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := repository.FishDestroy(id)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "削除しました")
}
