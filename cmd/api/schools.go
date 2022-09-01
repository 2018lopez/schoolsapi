//Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//create school handler for "POST" /v1/schools/ endpoint

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "New School")
}

// show school  GET /v1/schools/:id
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {

	//use the ParamsFromContext function to get the request contect as a slice
	params := httprouter.ParamsFromContext(r.Context())

	//get the value of the ":id" parameter

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	//display school id
	fmt.Fprintf(w, "Schools  %d\n", id)

}
