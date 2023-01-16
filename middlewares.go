package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := false
			fmt.Println("Checking Authentication")
			if flag {
				hf(w, r)
			} else {
				AuthenticationError(w, r)
			}
		}
	}
}

func LoggerMiddleware() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			method := r.Method
			fmt.Printf("%s | %s -> status \n", method, path)
			hf(w, r)
		}
	}
}
