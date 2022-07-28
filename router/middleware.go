package router

import (
	"anyweb/auth"
	"anyweb/util"
	"errors"
	"net/http"
	"strings"
)

const (
	AuthHeaderName   string = "Authorization"
	AuthHeaderPrefix        = "Bearer "
)

func ResponseMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(rw, r)
		if s := rw.Header().Get("err"); s != "" {
			rw.Write([]byte(s))
		}
	})
}

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		tokenString, err := extractToken(r.Header)

		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte(err.Error()))
			return
		}

		e, err := util.ParseJwtToken(tokenString)

		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte(err.Error()))
			return
		}

		ctx = auth.ContextWithMember(ctx, auth.NewMember(e, "", "", auth.NewGender(false)))

		r = r.WithContext(ctx)

		h.ServeHTTP(rw, r)
	})
}

func extractToken(header http.Header) (string, error) {
	tokenString := header.Get(AuthHeaderName)

	if tokenString == "" {
		return "", errors.New("please check your header " + AuthHeaderName)
	}

	return strings.TrimPrefix(tokenString, AuthHeaderPrefix), nil
}
