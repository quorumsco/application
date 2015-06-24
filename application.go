/*
Package application provides a convenient wrapper for common web
applications components.
*/
package application

import (
	"fmt"
	"net/http"

	"github.com/iogo-framework/logs"
)

// Application is an abstraction layer containing the router of your
// choice and additionnal components (templates, database links, etc).
type Application struct {
	Mux        Mux
	Components map[string]interface{}
}

// New creates a new Application with initialised fields.
func New() (*Application, error) {
	var app = new(Application)
	app.Components = make(map[string]interface{})
	return app, nil
}

// Get defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Get(path interface{}, handle http.HandlerFunc) {
	app.Mux.Get(path, handle)
}

// Post defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Post(path interface{}, handle http.HandlerFunc) {
	app.Mux.Post(path, handle)
}

// Put defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Put(path interface{}, handle http.HandlerFunc) {
	app.Mux.Put(path, handle)
}

// Patch defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Patch(path interface{}, handle http.HandlerFunc) {
	app.Mux.Patch(path, handle)
}

// Delete defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Delete(path interface{}, handle http.HandlerFunc) {
	app.Mux.Delete(path, handle)
}

// Options defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Options(path interface{}, handle http.HandlerFunc) {
	app.Mux.Options(path, handle)
}

// Use defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Use(handler func(http.Handler) http.Handler) {
	app.Mux.Use(handler)
}

// ServeHTTP allows Mux to start processing HTTP requests. Satisfies net/http.Handler.
func (app *Application) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	app.Mux.ServeHTTP(w, req)
}

// Serve start a listening server on port.
func (app *Application) Serve(host string, port int) {
	logs.Info("Listening on http://%s:%d", host, port)
	logs.Critical(http.ListenAndServe(fmt.Sprintf(":%d", port), app.Mux))
}
