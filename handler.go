package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Response struct {
	Status string `json:"status"`
}

func StatusHandler(ctx echo.Context) error {
	status, statusStr := app.Status()

	r := &Response{
		Status: statusStr,
	}

	return ctx.JSON(status, r)
}

func StartHandler(ctx echo.Context) error {
	app.Start()

	r := &Response{
		Status: "ok",
	}

	return ctx.JSON(http.StatusOK, r)
}

func StopHandler(ctx echo.Context) error {
	app.Stop()

	r := &Response{
		Status: "ok",
	}

	return ctx.JSON(http.StatusOK, r)
}
