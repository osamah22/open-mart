package main

import "net/http"

func (app *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	data := app.newDataTemplate(r)
	app.render(w, 200, "home.html", data)
}
