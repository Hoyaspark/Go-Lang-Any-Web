package router

import (
	"database/sql"
	"net/http"
)

func New(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello, world!"))
	}
}
