package middleware

import (
	"log"
	"net/http"
)

// LogReq middleware is used for logging incoming requests
func LogReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := r.URL.String()
		method := r.Method
		log.Println(method, uri)
		next.ServeHTTP(w, r)
	})
}
