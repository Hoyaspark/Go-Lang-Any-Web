package router

import (
	"anyweb/auth"
	"anyweb/config"
	"database/sql"
	"encoding/json"
	"net/http"
)

func LoginRoute(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		var param auth.LoginRequestBody

		if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
			panic(err)
		}

		res, err := auth.Login(config.ContextWithDatabase(r.Context(), db), &param)

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

func JoinRoute(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var param auth.JoinRequestBody

		err := json.NewDecoder(r.Body).Decode(&param)

		if err != nil {
			panic(err)
		}

		if err := auth.Join(config.ContextWithDatabase(r.Context(), db), &param); err != nil {
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
//		u, err := auth.UserFromContext(ctx)
//
//		if err != nil {
//			res.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//
//	}
//
//}
