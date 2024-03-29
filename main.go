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

	db := config.NewDatabase()

	defer db.Close()

	r := chi.NewRouter()

	apiRouter := chi.NewRouter()

	apiRouter.Use(middleware.DefaultLogger)
	apiRouter.Use(middleware.Recoverer)

	apiRouter.Group(func(r chi.Router) {
		r.Use(router.ResponseMiddleware)
		r.Post("/auth/login", router.LoginRoute(db))
		r.Post("/auth/register", router.JoinRoute(db))
	})

	apiRouter.Group(func(r chi.Router) {
		r.Use(router.AuthMiddleware)
		r.Use(router.ResponseMiddleware)
		r.Get("/my", router.MyPageRoute(db))
		r.Get("/board", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("Hello, Worlds!"))
		})
	})

	r.Mount("/api", apiRouter)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
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
