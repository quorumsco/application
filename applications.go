package applications

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/iogo-framework/databases"
	"github.com/iogo-framework/logs"
	"github.com/iogo-framework/router"
	"github.com/iogo-framework/settings"
	"github.com/iogo-framework/views"
)

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

type Application struct {
	Mux       Mux
	DB        *databases.DB
	Templates map[string]*template.Template
	Urls      map[string]string
}

func New() (*Application, error) {
	var app = new(Application)

	app.Urls = make(map[string]string)
	app.DB = new(databases.DB)
	app.Mux = router.New()

	funcMap := template.FuncMap{
		"path": func(name string, params ...interface{}) string {
			return fmt.Sprintf(app.Urls[name], params...)
		},
	}

	var err error
	if app.DB.SQLX, err = databases.InitSQLX(); err != nil {
		return nil, err
	}
	//if application.DB.GORM, err = databases.InitGORM(); err != nil {
	//return err
	//}
	app.Templates = views.Templates(&funcMap)

	return app, nil
}

func (app *Application) Name(urlFormat string, name string) {
	app.Urls[name] = urlFormat
}

func (app *Application) Path(name string, params ...interface{}) string {
	return fmt.Sprintf(app.Urls[name], params...)
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

func (app *Application) Serve() {
	logs.Info("Listening on http://localhost:%s", settings.Port)
	logs.Critical(http.ListenAndServe(":"+settings.Port, app.Mux))
}

func (app *Application) Load(urls func(app *Application)) {
	urls(app)
}

//func (app *Application) Get(pattern interface{}, controller interface{}, name string) {
//switch v := pattern.(type) {
//case string:
//app.Name(v, name)
//case regexp.Regexp:
//r := regexp.MustCompile("(\\(.+\\))")
//format := r.ReplaceAllLiteralString(v.String(), "%v")
//app.Name(format, name)
//fmt.Println(format)
//}

//goji.Get(pattern, controller)
//}
