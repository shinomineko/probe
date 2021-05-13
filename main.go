package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var app Application

func main() {
	app.Start()

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/status", StatusHandler)
	e.GET("/start", StartHandler)
	e.GET("/stop", StopHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
