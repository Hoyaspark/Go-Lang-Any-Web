package router

import (
	"anyweb/config"
	"anyweb/user"
	"database/sql"
	"encoding/json"
	"net/http"
)

func LoginRoute(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		var param user.LoginRequestBody

		if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
			panic(err)
		}

		res, err := user.Login(config.ContextWithDatabase(r.Context(), db), &param)

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

func Join(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var u user.User

		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			panic(err)
		}

		if err := user.Join(config.ContextWithDatabase(r.Context(), db), &u); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}

		rw.Write([]byte("Success"))

	}
}

//
//func MyPage(db *sql.DB) http.HandlerFunc {
//	return func(res http.ResponseWriter, req *http.Request) {
//		ctx := req.Context()
//
//		ctx = config.ContextWithDatabase(ctx, db)
//
//		req = req.WithContext(ctx)
//
//		u, err := user.UserFromContext(ctx)
//
//		if err != nil {
//			res.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//
//	}
//
//}
