package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var app Application

func main() {
	app.Start()

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/status", StatusHandler)
	e.POST("/start", StartHandler)
	e.POST("/stop", StopHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
