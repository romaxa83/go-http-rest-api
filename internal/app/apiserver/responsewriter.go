package apiserver

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	code int
}

// оставляем всю функциональность ResponseWriter, добавляем только статус код
func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}