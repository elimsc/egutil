package egu

import "net/http"

// https://gowebexamples.com/advanced-middleware/

// Middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// NewMiddleware create middleware
func NewMiddleware(f http.HandlerFunc) Middleware {
	middleware := func(next http.HandlerFunc) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			f(w, r)
			next(w, r)
		}

		return handler
	}

	return middleware
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
