package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	data := app.newDataTemplate(r)
	app.render(w, 200, "home.html", data)
}

func (app *Server) handleListPosts(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	categoryExists := app.categoryService.SlugExists(r.Context(), slug)

	if !categoryExists {
		app.notFound(w, r)
		return
	}

	data := app.newDataTemplate(r)
	app.render(w, 200, "posts.html", data)
}
