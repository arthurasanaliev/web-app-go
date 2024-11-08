package main

import (
	"github.com/arthurasanaliev/web-app-go/pkg/config"
	"github.com/arthurasanaliev/web-app-go/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

// routes defines handler functions
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(noSurf)
	mux.Use(sessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
