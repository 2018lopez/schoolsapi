//Filename: cmd/api/helpers.go

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// define a new type name envelope
type envelope map[string]interface{}

func (app *application) readIDParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())

	//get the value of the ":id" parameter

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil

}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	//convert our map into a JSON object
	js, err := json.MarshalIndent(data, " ", "\t")

	if err != nil {

		return err
	}

	//add a newline to make viewing on the terminal easier
	js = append(js, '\n')

	//add the headers
	for key, value := range headers {
		w.Header()[key] = value
	}

	//specify that we will serve our response using JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	//write the [] byte slice containing the JSON response body
	w.Write(js)

	return nil

}
