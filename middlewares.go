package applications

import (
	"net/http"

	"github.com/iogo-framework/router"
)

func (app *Application) ApplyTemplates(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		router.GetContext(r).Env["Templates"] = app.Templates
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (app *Application) ApplyDB(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		router.GetContext(r).Env["DB"] = app.DB
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
