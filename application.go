/*
Package application provides a convenient wrapper for common web
applications components.
*/
package application

import (
	"net/http"

	"github.com/quorumsco/gojimux"
	"github.com/quorumsco/logs"
)

// Application is an abstraction layer containing the router of your
// choice and additionnal components (templates, database links, etc).
type Application struct {
	Components map[string]interface{}
}

// New creates a new Application with initialised fields.
func New() *Application {
	var app = new(Application)
	app.Components = make(map[string]interface{})
	return app
}

// Get defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Get(path interface{}, handle http.HandlerFunc) {
	gojimux.Get(path, handle, app.Components["Mux"])
}

// Post defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Post(path interface{}, handle http.HandlerFunc) {
	gojimux.Post(path, handle, app.Components["Mux"])
}

// Put defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Put(path interface{}, handle http.HandlerFunc) {
	gojimux.Put(path, handle, app.Components["Mux"])
}

// Patch defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Patch(path interface{}, handle http.HandlerFunc) {
	gojimux.Patch(path, handle, app.Components["Mux"])
}

// Delete defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Delete(path interface{}, handle http.HandlerFunc) {
	gojimux.Delete(path, handle, app.Components["Mux"])
}

// Options defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Options(path interface{}, handle http.HandlerFunc) {
	gojimux.Options(path, handle, app.Components["Mux"])
}

// Use defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Use(handler func(http.Handler) http.Handler) {
	gojimux.Use(handler, app.Components["Mux"])
}

// ServeHTTP allows Mux to start processing HTTP requests. Satisfies net/http.Handler.
func (app *Application) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	gojimux.ServeHTTP(w, req, app.Components["Mux"])
}

// Serve start a listening server on port.
func (app *Application) Serve(listen string) error {
	logs.Info("listening on http://%s", listen)
	return ListenAndServe(listen, app.Components["Mux"])
	// return gojimux.ListenAndServe(listen, app.Components["Mux"])
}
