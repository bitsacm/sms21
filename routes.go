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
		"/auth/google",
		withMiddleware(
			[]middleware.Middleware{
				middleware.LogReq,
			},
			api.GoogleAuthScreen(env.models),
		).ServeHTTP,
	)

	r.HandleFunc(
		"/auth/google/redirect",
		withMiddleware(
			[]middleware.Middleware{
				middleware.LogReq,
			},
			api.GoogleTokenExchange(env.models),
		).ServeHTTP,
	)

	r.HandleFunc(
		"/stock",
		withMiddleware(
			[]middleware.Middleware{
				middleware.LogReq,
			},
			api.AddStockHandler(env.models),
		).ServeHTTP,
	)

	r.HandleFunc(
		"/stock/{id}",
		withMiddleware(
			[]middleware.Middleware{
				middleware.LogReq,
			},
			api.GetStockHandler(env.models),
		).ServeHTTP,
	)

	r.HandleFunc(
		"/transaction",
		withMiddleware(
			[]middleware.Middleware{
				middleware.LogReq,
			},
			api.AddTransactionHandler(env.models),
		).ServeHTTP,
	)

	return r
}
