package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func ()  {
		if err := recover(); err != nil {
			w.Header().Set("Connection", "close")
			app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
		}
	}()
	next.ServeHTTP(w,r)
	})
} 
func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			proto  = r.Proto
			uri    = r.URL.RequestURI()
			method = r.Method
		)
		app.logger.Info("request recieved", "ip", ip, "proto", proto, "uri", uri, "method", method)

		next.ServeHTTP(w,r)
	})
}