package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
}

func StatusHandler(ctx echo.Context) error {
	status, statusMsg := app.Status()
	if status != http.StatusOK {
		return echo.NewHTTPError(status, statusMsg)
	}
	r := &Response{
		Message: statusMsg,
	}
	return ctx.JSON(status, r)
}

func StartHandler(ctx echo.Context) error {
	app.Start()
	r := &Response{
		Message: "ok",
	}
	return ctx.JSON(http.StatusOK, r)
}

func StopHandler(ctx echo.Context) error {
	app.Stop()
	r := &Response{
		Message: "ok",
	}
	return ctx.JSON(http.StatusOK, r)
}
