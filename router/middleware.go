package router

import (
	"anyweb/user"
	"anyweb/util"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		uuid := uuid.NewString()

		logger := util.LoggerFunc(func(message string, err error) {
			if err != nil {
				fmt.Printf("[Error][Request][%s][%s] : %s\n", uuid, time.Now().Format(time.RFC3339), err.Error())
				return
			}
			fmt.Printf("[Normal][Request][%s][%s] : %s\n", uuid, time.Now().Format(time.RFC3339), message)
		})

		logger.Log("connect", nil)

		//ctx = context.WithValue(ctx, config.Logger, logger)

		r = r.WithContext(ctx)

		h.ServeHTTP(rw, r)

	})
}

func ResponseMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(rw, r)
	})
}

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		u, err := util.ParseJwtToken(tokenString)

		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte(err.Error()))
			return
		}

		ctx = user.ContextWithUser(ctx, u)

		r = r.WithContext(ctx)

		h.ServeHTTP(rw, r)
	})
}
