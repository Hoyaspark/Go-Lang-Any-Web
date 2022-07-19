package main

import (
	"anyweb/config"
	"anyweb/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func main() {

	db := config.NewMySQL()

	defer db.Close()

	apiRouter := chi.NewRouter()

	r := chi.NewRouter()

	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)

	r.Group(func(r chi.Router) {
		r.Get("/auth/login", router.New(db))
	})

	r.Group(func(r chi.Router) {
		r.Use(router.AuthMiddleware)

		r.Get("/board", router.New(db))
	})

	apiRouter.Mount("/api", r)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      apiRouter,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	err := s.ListenAndServe()

	defer s.Close()

	if err != nil {
		panic(err)
	}
}
