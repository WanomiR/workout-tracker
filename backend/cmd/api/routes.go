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
	mux.Post("/register", nil)
	mux.Get("/refresh", nil)
	mux.Get("/logout", nil)

	mux.Get("/", app.Home)

	mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+app.Port+"/swagger/doc.json"),
	))

	return mux
}
