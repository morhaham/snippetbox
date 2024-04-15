package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.HandleFunc("GET /{$}", app.home)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /snippet/{id}", app.showSnippet)
	mux.HandleFunc("POST /snippet/create", app.createSnippet)
	mux.HandleFunc("GET /snippet/create", app.createSnippetForm)
	return standardMiddleware.Then(mux)
}
