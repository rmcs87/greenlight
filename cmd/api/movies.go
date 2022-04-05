package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new movie\n")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Show details for movie %d\n", id)
}
