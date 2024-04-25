package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /{$}", dynamicMiddleware.ThenFunc(app.home))
	mux.HandleFunc("GET /snippet/{id}", app.showSnippet)
	mux.HandleFunc("POST /snippet/create", app.createSnippet)
	mux.HandleFunc("GET /snippet/create", app.createSnippetForm)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
