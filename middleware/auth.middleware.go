package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dush-t/sms21/db/models"
)

// Middleware is any function that accepts an http handler function
// and returns a new http handler function
type Middleware func(http.Handler) http.Handler

// Auth will add the User struct making the request, to the
// request context
func Auth(data models.Models) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := strings.Split(r.Header.Get("Authorization"), " ")[1]
			claims := &models.Claims{}

			tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte("lolmao12345"), nil
			})
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if !tkn.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			user, err := data.Users.GetByID(claims.ID)
			key := "user"
			ctx := context.WithValue(r.Context(), key, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
