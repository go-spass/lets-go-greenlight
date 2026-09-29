package main

import (
	"encoding/json/v2" // New import
	"net/http"
)

// Declare a handler which writes a plain-text response with information about
// the application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("entered healthcheckHandler", "method", r.Method, "request", r.RequestURI)

	// Create a map which holds the information that we want to send in the response.
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// Pass the map to the json.Marshal() function, without specifying any non-default
	// options. This returns a []byte slice containing the encoded JSON. If there was
	// an error, we log it and send the client a generic error message. Use the
	// json.Deterministic(true) option to force the data map to be encoded in a
	// predictable, deterministic way.
	js, err := json.Marshal(data, json.Deterministic(true))
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request",
			http.StatusInternalServerError)
		return
	}

	// Append a newline to the JSON. This is just a small nicety to make it easier to
	// view in terminal applications.
	js = append(js, '\n')

	// Set the "Content-Type: application/json" header on the response. If you forget to do
	// this, Go will default to sending a "Content-Type: text/plain; charset=utf-8"
	// header instead.
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON as the HTTP response body.
	_, _ = w.Write([]byte(js))
}
