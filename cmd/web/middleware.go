package main

import (
	"fmt"
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// split for readability
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		// having return before calling next, the chain will stop and go back up
		// above is code executed down the chain
		next.ServeHTTP(w, r)
		// below is code executed up back the chain
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// a deferred func which will always be run in the event of a panic as Go unwinds the stack
		defer func() {
			// use the builtin recover to check if theres been a panic
			// this func will catch panics that happen in the same gorouting that executed the recoverPanic
			// if a handler spins up other goroutins then any panics will not be recovered
			if err := recover(); err != nil {
				// with this header set Go's http server auto closes the cur conn after a res has been sent
				// also informs that user the conn will be closed
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
