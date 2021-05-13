package main

import "net/http"

type Application struct {
	isReady bool
}

func (app *Application) Start() {
	app.isReady = true
}

func (app *Application) Stop() {
	app.isReady = false
}

func (app *Application) Status() (int, string) {
	if app.isReady {
		return http.StatusOK, "up"
	}
	return http.StatusServiceUnavailable, "down"
}
