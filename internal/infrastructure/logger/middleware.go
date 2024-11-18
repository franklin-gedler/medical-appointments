package logger

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// LoggerMiddleware es un middleware que loguea las solicitudes HTTP usando zap.Logger.
func LoggerMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK, size: 0}
			next.ServeHTTP(ww, r)
			duration := time.Since(start)

			logger.Debug("HTTP request",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
				zap.String("remote_addr", r.RemoteAddr),
				zap.Int("status", ww.statusCode),
				zap.Int64("size", ww.size),
				zap.Duration("duration", duration),
			)
		})
	}
}

// responseWriter es un http.ResponseWriter que captura el código de estado y el tamaño de la respuesta.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int64
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += int64(size)
	return size, err
}
