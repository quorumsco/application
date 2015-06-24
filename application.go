/*
Package application provides a convenient wrapper for common web
applications components.
*/
package application

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/iogo-framework/logs"
)

// Application an abstraction wrapper containing the router of your
// choice, a database link and some templates.
type Application struct {
	Mux       Mux
	DB        DB
	Templates map[string]*template.Template
}

// New creates a new Application with initialised fields.
func New() (*Application, error) {
	var app = new(Application)

	app.Templates = make(map[string]*template.Template)

	return app, nil
}

// Get defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Get(path string, handle http.HandlerFunc) {
	app.Mux.Get(path, handle)
}

// Post defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Post(path string, handle http.HandlerFunc) {
	app.Mux.Post(path, handle)
}

// Put defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Put(path string, handle http.HandlerFunc) {
	app.Mux.Put(path, handle)
}

// Patch defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Patch(path string, handle http.HandlerFunc) {
	app.Mux.Patch(path, handle)
}

// Delete defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Delete(path string, handle http.HandlerFunc) {
	app.Mux.Delete(path, handle)
}

// Options defers to the router Mux the handling of requests matching the
// pattern path, associated with the handle function.
func (app *Application) Options(path string, handle http.HandlerFunc) {
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
func (app *Application) Serve(port int) {
	logs.Info("Listening on http://localhost:%d", port)
	logs.Critical(http.ListenAndServe(fmt.Sprintf(":%d", port), app.Mux))
}
