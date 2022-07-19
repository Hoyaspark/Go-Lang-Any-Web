package route

import (
	"database/sql"
	"net/http"
)

type AuthRouter *http.ServeMux

func New(db *sql.DB) AuthRouter {
	mux := http.NewServeMux()

	return AuthRouter(mux)
}
