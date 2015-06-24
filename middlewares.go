package application

import (
	"net/http"

	"github.com/iogo-framework/router"
)

// Apply is a middleware which make the app context available to the
// requests handlers.
func (app *Application) Apply(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		router.GetContext(r).Env["Templates"] = app
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
