package httpserver

import (
	"github.com/dingowd/CyberZilla/test3/internal/logger"
	"net/http"
)

func loggingMiddleware(f http.HandlerFunc, logg logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := r.Method + " " + r.RequestURI
		logg.Info(s)
		f(w, r)
	}
}
