package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := &ResponseWrapper{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		start := time.Now()
		next.ServeHTTP(writer, r)
		log.Println("[REQUEST]", r.Method, "Status:", writer.StatusCode, "| Endpoint:", r.URL.Path, "| Duration:", time.Since(start).Microseconds(), "ms")

	})
}
