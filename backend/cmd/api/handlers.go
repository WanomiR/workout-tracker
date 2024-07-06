package main

import "net/http"

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		"active", "Workout Tracker is up and running", "0.0.0",
	}

	writeJSONResponse(w, http.StatusOK, payload)
}
