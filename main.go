package main

import (
	"log"
	"os"

	"go_rest_api/handler"
	"go_rest_api/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var e = createMux()

func main() {
	port := os.Getenv("PORT")
	db := connectDB()
	repository.SetDB(db)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/places", handler.PlaceIndex)

	e.Logger.Fatal(e.Start(":" + port))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())

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
