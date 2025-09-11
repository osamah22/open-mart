package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

func (app *Server) newDataTemplate(r *http.Request) map[string]any {
	data := map[string]any{}
	if app.sessionManager.Exists(r.Context(), "flash") {
		fmt.Print("getting flash message")
		data["flash"] = app.sessionManager.Pop(r.Context(), "flash")
	}
	data["current_year"] = time.Now().Year()

	return data
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Get all page templates
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Expand the partials glob
		partials, err := filepath.Glob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// Build full pattern list: base, partials..., page
		patterns := append([]string{"./ui/html/base.html"}, partials...)
		patterns = append(patterns, page)

		// Parse all templates together
		ts, err := template.New(name).ParseFiles(patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
