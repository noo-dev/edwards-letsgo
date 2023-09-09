package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) initRoutes() http.Handler {

	// create a standart middleware chain which
	// will be used for every request
	standartMiddleware := alice.New(app.recoverPanic, app.logRequestMw, secureHeadersMw)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.Home))
	mux.Get("/snippet/create", http.HandlerFunc(app.CreateSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.CreateSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.ShowSnippet))
	fileServer := http.FileServer(http.Dir("../../ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standartMiddleware.Then(mux)
}
