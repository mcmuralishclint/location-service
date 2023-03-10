package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Now()
		log.WithFields(log.Fields{
			"method":     r.Method,
			"requestURI": r.RequestURI,
			"remoteAddr": r.RemoteAddr,
			"duration":   end.Sub(start),
		}).Info("HTTP Request")
	})
}
