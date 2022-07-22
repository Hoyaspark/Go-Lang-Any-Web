package main

import (
	"anyweb/config"
	"anyweb/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {

	db := config.NewDatabase()

	defer db.Close()

	r := chi.NewRouter()

	apiRouter := chi.NewRouter()

	apiRouter.Use(middleware.DefaultLogger)
	apiRouter.Use(middleware.Recoverer)
	apiRouter.Use(router.ResponseMiddleware)

	apiRouter.Group(func(r chi.Router) {
		r.Post("/auth/login", router.LoginRoute(db))
		r.Post("/auth/register", router.Join(db))
		r.Get("/test", func(rw http.ResponseWriter, r *http.Request) {
			log.Println(runtime.NumGoroutine())

			time.Sleep(time.Second * 5)
			rw.Write([]byte("Hello, Worlds!"))
		})
	})

	apiRouter.Group(func(r chi.Router) {
		r.Use(router.AuthMiddleware)
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
