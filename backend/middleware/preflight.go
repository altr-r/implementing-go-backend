package middleware

import "net/http"

func Preflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handle preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
		}
		next.ServeHTTP(w, r)
	})
}
