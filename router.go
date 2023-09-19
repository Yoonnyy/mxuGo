package main

import (
	"net/http"

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

	router.Mount("/api/v1", v1Routes())

	return router
}

func v1Routes() http.Handler {
	router := chi.NewRouter()

	router.Get("/{slug}", controllers.SearchSlug)
	router.
		With(middleware.AllowContentType("multipart/form-data")).
		Post("/", controllers.CreateShortened)
	router.Get("/s", func(w http.ResponseWriter, r *http.Request) {
		panic("AAAAAAAAAAA")
	})

	return router
}
