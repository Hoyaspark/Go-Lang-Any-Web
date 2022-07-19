package router

import (
	"anyweb/config"
	"anyweb/user"
	"database/sql"
	"encoding/json"
	"net/http"
)

func Login(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		var u user.User

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			panic(err)
		}

		res, err := user.Login(config.ContextWithDatabase(r.Context(), db), &u)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}

		err = json.NewEncoder(rw).Encode(&res)

		if err != nil {
			panic(err)
		}
	}
}

func Join() {

}
