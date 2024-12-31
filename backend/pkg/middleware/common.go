package middleware

import "net/http"

type ResponseWrapper struct {
	http.ResponseWriter
	StatusCode int
}

func (w *ResponseWrapper) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}
