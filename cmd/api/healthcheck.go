package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, 200, envelope{"health": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
