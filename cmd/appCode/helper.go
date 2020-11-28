package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/CS-PCockrill/queue/pkg/models"
)

func (app *appInjection) isAuthenticated(user models.User) bool {
	// Authenticated whether the user is logged in
	// Return T or F
	return false
}

// The serverError helper writes an error message and stack trace to the errorLog,
//then sends a generic 500 Internal Server Error response to the user.
func (app *appInjection) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	_ = app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
//to the user.
func (app *appInjection) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *appInjection) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
