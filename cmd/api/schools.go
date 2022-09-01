//Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
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

	//display school id
	fmt.Fprintf(w, "Schools  %d\n", id)

}
