//Filename - cmd/api/errros.go

package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {

	app.logger.Println(err)
}

// we want to send JSON formate error message
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	//create json response
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Server error resonse
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//log the error
	app.logError(r, err)
	//prepare msg with the error
	message := "the server encounted a problem and could not process the request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)

}

// Not Found Response

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	//create msg
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// Method not allowed response
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	//create msg
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
