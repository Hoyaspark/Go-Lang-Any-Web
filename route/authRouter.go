package route

import (
	"database/sql"
	"net/http"
)

func New(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

	})
}
