package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about
// the application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("entered healthcheckHandler", "method", r.Method, "request", r.RequestURI)

	_, _ = fmt.Fprintln(w, "status: available")
	_, _ = fmt.Fprintf(w, "environment: %s\n", app.config.env)
	_, _ = fmt.Fprintf(w, "version: %s\n", version)
}
