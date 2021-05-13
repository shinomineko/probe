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

func (app *Application) Hello() int {
	if app.isReady == true {
		return http.StatusOK
	}
	return http.StatusServiceUnavailable
}

func (app *Application) Status() string {
	var status string
	if app.isReady == true {
		status = "up"
	} else {
		status = "down"
	}
	return status
}
