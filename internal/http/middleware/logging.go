package middleware

import (
	"kasir-api/internal/pkg"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// ResponseWriter wrapper untuk menangkap status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.statusCode = code
		rw.written = true
		rw.ResponseWriter.WriteHeader(code)
	}
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.written {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(b)
}

// LoggingMiddleware mencatat setiap request yang masuk dengan structured logging
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)

		pkg.Log.WithFields(logrus.Fields{
			"method":      r.Method,
			"path":        r.URL.Path,
			"status_code": wrapped.statusCode,
			"duration_ms": duration.Milliseconds(),
		}).Info("request completed")
	})
}
