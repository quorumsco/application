package application

import (
	"net/http"

	"github.com/silverwyrda/iogo"
)

func (app *Application) ApplyTemplates(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		iogo.GetContext(r).Env["Templates"] = app.Templates
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (app *Application) ApplyDB(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		iogo.GetContext(r).Env["DB"] = app.DB
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
