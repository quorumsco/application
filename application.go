package application

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/iogo-framework/logs"
)

type Application struct {
	Mux       Mux
	DB        DB
	Templates map[string]*template.Template
}

func New() (*Application, error) {
	var app = new(Application)

	app.Templates = make(map[string]*template.Template)

	return app, nil
}

func (app *Application) Get(path string, handle http.HandlerFunc) {
	app.Mux.Get(path, handle)
}

func (app *Application) Post(path string, handle http.HandlerFunc) {
	app.Mux.Post(path, handle)
}

func (app *Application) Put(path string, handle http.HandlerFunc) {
	app.Mux.Put(path, handle)
}

func (app *Application) Patch(path string, handle http.HandlerFunc) {
	app.Mux.Patch(path, handle)
}

func (app *Application) Delete(path string, handle http.HandlerFunc) {
	app.Mux.Delete(path, handle)
}

func (app *Application) Options(path string, handle http.HandlerFunc) {
	app.Mux.Options(path, handle)
}

func (app *Application) Use(handler func(http.Handler) http.Handler) {
	app.Mux.Use(handler)
}

func (app *Application) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	app.Mux.ServeHTTP(w, req)
}

func (app *Application) Serve(port int) {
	logs.Info("Listening on http://localhost:%d", port)
	logs.Critical(http.ListenAndServe(fmt.Sprintf(":%d", port), app.Mux))
}
