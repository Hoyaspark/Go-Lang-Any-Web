package middleware

import (
	"anyweb/config"
	"anyweb/util"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
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

		ctx = context.WithValue(ctx, config.Logger, logger)

		r = r.WithContext(ctx)

		h.ServeHTTP(rw, r)

	})
}

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		logger := ctx.Value(config.Logger).(util.Logger)

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			logger.Log("Not Authorized", nil)
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return config.AuthProperties.JwtSecret, nil
		})

		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			logger.Log("", err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			email := claims["userEmail"].(string)
			ctx = context.WithValue(ctx, config.JWTInfo, email)
			r = r.WithContext(ctx)
		}

		h.ServeHTTP(rw, r)
	})
}
