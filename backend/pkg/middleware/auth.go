package middleware

import (
	"net/http"

	"github.com/Vitaljano/invly/backend/internal/auth"
)

func AuthMiddleware(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, auth.ErrorMissingAuthorization, http.StatusUnauthorized)
				return
			}

		})
	}
}
