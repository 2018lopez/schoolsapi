//Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"schools.ImerLopez.net/internal/data"
)

//create school handler for "POST" /v1/schools/ endpoint

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "New School")
}

// show school  GET /v1/schools/:id
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	//create a new instance of the school struct containing the id we extracted from our url and some sample data

	school := data.School{
		ID:       id,
		CreateAt: time.Now(),
		Name:     "Apple Tree",
		Level:    "High School",
		Contact:  "Anna Smith",
		Phone:    "601-4412",
		Address:  "14 Apple Street",
		Mode:     []string{"blended", "online"},
		Version:  1,
	}

	err = app.writeJSON(w, http.StatusOK, school, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could processs your request", http.StatusInternalServerError)
	}
}
