package main

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/handlers"
	"github.com/go-chi/chi"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
