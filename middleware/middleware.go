package middleware

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rs/cors"

	"cinemanz/helper/response"
)

const (
	tokenName = "_cnz"
	secretKey = "c0b4d1b4c4"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct{}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS() *cors.Cors {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	return cors
}

// IsAuthenticated will the AUthenticated middleware
func (m *GoMiddleware) IsAuthenticated(next http.Handler) http.Handler {
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

// InitMiddleware intialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
