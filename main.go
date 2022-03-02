package main

import (
	"log"
	"os"

	_ "go_rest_api/docs"
	"go_rest_api/handler"
	"go_rest_api/repository"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var e = createMux()

// @title GO_REST_API
// @version 0.1
// @description Go REST API
func main() {
	port := os.Getenv("PORT")
	db := connectDB()
	repository.SetDB(db)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/fishes", handler.FishIndex)
	e.POST("/fishes", handler.FishCreate)
	e.GET("/fishes/:id", handler.FishShow)
	e.PATCH("/fishes/:id", handler.FishUpdate)
	e.DELETE("/fishes/:id", handler.FishDestroy)

	// Validate メソッドを使えるようにする
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Logger.Fatal(e.Start(":" + port))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}

	log.Println("db connection succeeded")

	return db
}

type CustomValidator struct {
	validator *validator.Validate
}

// 自作の Validate メソッドを用意、その中で validator ライブラリを使用する
func (cv *CustomValidator) Validate(i interface{}) error {
  return cv.validator.Struct(i)
}
