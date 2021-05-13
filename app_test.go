package main

import "testing"

func TestReadiness(t *testing.T) {
	assertReadiness := func(t testing.TB, app Application, want bool) {
		t.Helper()
		got := app.isReady

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("start", func(t *testing.T) {
		app := Application{}
		app.Start()
		assertReadiness(t, app, true)
	})

	t.Run("stop", func(t *testing.T) {
		app := Application{}
		app.Stop()
		assertReadiness(t, app, false)
	})
}

func TestStatus(t *testing.T) {
	assertStatus := func(t testing.TB, app Application, want int) {
		t.Helper()
		got, _ := app.Status()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("start", func(t *testing.T) {
		app := Application{}
		app.Start()
		assertStatus(t, app, 200)
	})

	t.Run("stop", func(t *testing.T) {
		app := Application{}
		app.Stop()
		assertStatus(t, app, 503)
	})
}
