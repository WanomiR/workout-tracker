package main

import (
	_ "backend/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func (app *App) Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Post("/authenticate", app.Authenticate)
	mux.Post("/register", app.RegisterUser)
	mux.Get("/refresh", app.RefreshToken)
	mux.Get("/logout", app.Logout)

	// TODO:
	//  - protect all other routes
	//  - add StatusUnauthorized to all handlers' swagger
	mux.Route("/tracker", func(mux chi.Router) {
		mux.Use(app.requireAuthentication)

	})

	mux.Get("/", app.Home)

	mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+app.Port+"/swagger/doc.json"),
	))

	return mux
}
