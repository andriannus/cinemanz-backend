package auth

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"cinemanz/helper/response"
)

const (
	tokenName = "_cnz"
	secretKey = "c0b4d1b4c4"
)

// IsAuthenticated will check token
func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cnzToken := r.Header.Get(tokenName)

		if cnzToken == "" {
			response.Error(w, http.StatusUnauthorized, "Token not found")
		}

		token, err := jwt.Parse(cnzToken, func(token *jwt.Token) (interface{}, error) {
			method, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok || method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}
