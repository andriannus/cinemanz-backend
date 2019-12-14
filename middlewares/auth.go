package middlewares

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"cinemanz/utils"
)

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

		if err != nil || !token.Valid {
			utils.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}
