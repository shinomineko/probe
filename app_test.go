package main

import "testing"

func TestStart(t *testing.T) {
	app := Application{}
	app.Start()
	got := app.isReady
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestStop(t *testing.T) {
	app := Application{}
	app.Stop()
	got := app.isReady
	want := false
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("start", func(t *testing.T) {
		app := Application{}
		app.Start()
		got := app.Hello()
		want := 200
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("stop", func(t *testing.T) {
		app := Application{}
		app.Stop()
		got := app.Hello()
		want := 503
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
