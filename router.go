package main

import (
	"net/http"
	"strings"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func ExtractBasePath(path string) string {
	base_path := strings.Split(path, "/")[1]
	path = strings.Replace(path, "/"+base_path+"", "", 1)
	return path
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	path = ExtractBasePath(path)
	handler, exist := r.rules[path]
	return handler, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exist := r.FindHandler(request.URL.Path)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w, request)
}
