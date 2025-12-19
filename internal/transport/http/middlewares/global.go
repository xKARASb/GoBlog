package middlewares

import (
	"log/slog"
	"net/http"
)

func JSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{
			w, 200,
		}

		next.ServeHTTP(rw, r)
		slog.Info("http request", slog.String("method", r.Method), slog.String("endpoint", r.URL.Path), slog.Int("status", rw.statusCode))
	})
}
