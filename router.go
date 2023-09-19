package main

import (
	"github.com/Yoonnyy/GoMxu/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		// TODO: Allowed headers
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           60,
	}))

	router.Get("/{slug}", controllers.SearchSlug)
	router.
		With(middleware.AllowContentType("multipart/form-data")).
		Post("/", controllers.CreateShortened)

	return router
}
