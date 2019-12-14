package middlewares

import (
	"context"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"cinemanz/utils"
)

type key string

const contextKey key = "userInfo"

// IsAuthenticated check token
func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cnzToken := r.Header.Get("_cnz")

		if cnzToken == "" {
			utils.ResponseError(w, http.StatusUnauthorized, "Token not found")
			return
		}

		token, err := jwt.Parse(cnzToken, func(token *jwt.Token) (interface{}, error) {
			method, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return []byte("c0b4d1b4c4"), nil
		})

		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			utils.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		ctx := context.WithValue(context.Background(), contextKey, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
