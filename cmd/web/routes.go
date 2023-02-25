package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jamshu/go_web_app/pkg/config"
	"github.com/jamshu/go_web_app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	return mux
}
