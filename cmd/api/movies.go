package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Rixbur/greenlight/internal/data"
)

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movies{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablance",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// using anonymous structs
	// struct fields have to be exported
	// that is required by go's standard library package `reflect`
	var input struct {
		Title   string   `json"title"`
		Year    int32    `json"year"`
		Runtime int32    `json"runtime"`
		Genres  []string `json"genres`
	}

	// Any JSON key/value pair which cannot be successfully assigned
	// to the corresponding field in the input struct will be silently ignored.
	// there is no need to close r.Body, it is done by the http server
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input) // + includes the field name
}
