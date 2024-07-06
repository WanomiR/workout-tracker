package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	Domain string
	Port   string
	DSN    string // data source name
	DB     repository.DatabaseRepo
	Auth   *Auth
}

// @title Workout Tracker
// @version 0.0.0
// @description Service for keeping track of your training progress.

// @host localhost:8888
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	app := new(App)

	app.Domain = os.Getenv("DOMAIN")
	app.Port = os.Getenv("PORT")
	app.DSN = os.Getenv("DSN")

	app.Auth = &Auth{
		Issuer:        app.Domain,
		Secret:        os.Getenv("JWT_SECRET"),
		Audience:      app.Domain,
		TokenExpiry:   15 * time.Minute,
		RefreshExpiry: 24 * time.Minute,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  app.Domain,
	}

	conn, err := app.connectToDb()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDbRepo{Conn: conn}
	defer app.DB.Connection().Close()

	log.Fatal(http.ListenAndServe(":"+app.Port, app.Routes()))
}
