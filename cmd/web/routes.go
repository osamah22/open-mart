package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Server) routes() http.Handler {
	router := chi.NewRouter()

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w, r)
	})

	router.Use(app.recoverPanic)
	router.Use(app.logRequest)
	router.Use(secureHeaders)
	router.Use(app.sessionManager.LoadAndSave)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	router.Get("/", app.handleHome)
	router.Get("/{slug}", app.handleListPosts)
	return router
}
