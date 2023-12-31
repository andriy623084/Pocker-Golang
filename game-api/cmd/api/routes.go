package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *AppConfig) routes() http.Handler {
	mux := chi.NewRouter()
	// who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With", "X-HTTP-Method-Override"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/start-game", app.NewGame)
	mux.Get("/open-preflop", app.OpenPreFlop)


	return mux
}