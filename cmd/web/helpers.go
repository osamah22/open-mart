package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// render Retrieve the appropriate template set from the cache based on the page
// name (like 'home.html'). If no entry exists in the cache with the
// provided name, then create a new error and call the serverError() helper
// method that we made earlier and return.
func (app *Server) render(w http.ResponseWriter, status int, page string, data map[string]any) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *Server) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	app.render(w, 500, "server-error.html", nil)
}

func (app *Server) clientError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), http.StatusInternalServerError)
}

func (app *Server) notFound(w http.ResponseWriter) {
	app.render(w, 404, "not-found.html", nil)
}
