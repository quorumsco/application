package application

import "net/http"

type Mux interface {
	Get(path string, handle http.HandlerFunc)
	Post(path string, handle http.HandlerFunc)
	Put(path string, handle http.HandlerFunc)
	Patch(path string, handle http.HandlerFunc)
	Delete(path string, handle http.HandlerFunc)
	Options(path string, handle http.HandlerFunc)
	Use(handler func(http.Handler) http.Handler)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
