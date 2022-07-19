package main

import (
	"anyweb/config"
	"anyweb/middleware"
	"anyweb/route"
	"net/http"
	"time"
)

func main() {

	db := config.NewMySQL()

	defer db.Close()

	mux := http.NewServeMux()

	mux.Handle("/auth", middleware.LoggingMiddleware(route.New(db)))
	mux.Handle("/board", middleware.LoggingMiddleware(middleware.AuthMiddleware(route.New(db))))

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
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
