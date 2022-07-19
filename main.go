package main

import (
	"anyweb/config"
	"net/http"
	"time"
)

func main() {

	db := config.NewMySQL()

	defer db.Close()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
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
