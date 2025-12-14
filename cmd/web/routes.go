package main

// func (app *Server) routes() http.Handler {
// 	router := chi.NewRouter()
//
// 	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
// 		app.notFound(w, r)
// 	})
//
// 	router.Use(app.recoverPanic)
// 	router.Use(app.logRequest)
// 	router.Use(secureHeaders)
// 	router.Use(app.sessionManager.LoadAndSave) // for sessions managements
//
// 	fileServer := http.FileServer(http.Dir("./ui/static"))
// 	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
//
// 	// ---- Home ----
// 	router.Get("/", app.handleHome)
//
// 	// ---- Auth ----
// 	router.Get("/register", app.showSignUpForm)
// 	router.Post("/register", app.handleSignUp)
// 	router.Get("/login", app.showLoginForm)
// 	router.Post("/login", app.handleLogin)
//
// 	// ---- Categories ----
// 	router.Get("/{slug}", app.handleListPosts)
// 	return router
// }
