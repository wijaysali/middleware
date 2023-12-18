package main

import "net/http"

const USERNAME = "batman"
const PASSWORD = "secret"

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		username, password, ok := r.BasicAuth()
		if !ok{
			w.Write([]byte(`something went wrong`))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte(`username/password wrong`))
			return
		}
		next.ServeHTTP(w,r)
	})
}

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Method != "GET"{
			w.Write([]byte("Only GET IS ALLOWED"))
			return
		}
		next.ServeHTTP(w, r)
	})
}