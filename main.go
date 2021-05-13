package main

import (
	"github.com/labstack/echo"
)

var app Application

func main() {
	app.Start()

	e := echo.New()
	e.GET("/hello", HelloHandler)
	e.GET("/start", StartHandler)
	e.GET("/stop", StopHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
