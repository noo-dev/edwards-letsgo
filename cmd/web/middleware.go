package main

import (
	"fmt"
	"net/http"
)

func secureHeadersMw(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// add headers to response header
		w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
		w.Header().Set("X-Frame-Options", "deny")

		// call the next middleware/handler in the chain
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (app *application) logRequestMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log information about request
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recover panic
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func customHeaderMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Custom-Header", "Home page")
		next.ServeHTTP(w, r)
	})
}
