package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dush-t/sms21/api"
	"github.com/dush-t/sms21/middleware"
)

func withMiddleware(m []middleware.Middleware, h http.Handler) http.Handler {
	handler := h
	for i := len(m) - 1; i >= 0; i-- {
		mware := m[i]
		handler = mware(handler)
	}

	return handler
}

// Router will set up all the routes of the app
func Router(env Env) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc(
		"/sign_up",
		withMiddleware(
			[]middleware.Middleware{
				middleware.LogReq,
			},
			api.SignUpHandler(env.models),
		).ServeHTTP,
	)

	r.HandleFunc(
		"/sign_in",
		withMiddleware(
			[]middleware.Middleware{
				middleware.LogReq,
			},
			api.SignInHandler(env.models),
		).ServeHTTP,
	)

	return r
}
