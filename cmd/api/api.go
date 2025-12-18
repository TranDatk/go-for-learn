package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	postTransport "social/internal/posts/transport"
	userTransport "social/internal/users/transport"
)

type application struct {
	config  config
	handler handler
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type handler struct {
	userHandler *userTransport.UserHandler
	postHandler *postTransport.PostHandler
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", app.handler.userHandler.Register)
			r.Post("/", app.handler.userHandler.Register)
		})

		r.Route("/posts", func(r chi.Router) {
			r.Get("/{id}", app.handler.postHandler.NewPostService)
			r.Post("/", app.handler.postHandler.NewPostService)
		})
	})

	return r
}

func (app *application) run(handler http.Handler) error {

	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server started on %s", app.config.addr)

	return srv.ListenAndServe()
}
