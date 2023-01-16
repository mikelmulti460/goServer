package main

import (
	"net/http"
)

type Routers struct {
	api_router  *Router
	home_router *Router
}

type Server struct {
	port    string
	routers *Routers
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
		routers: &Routers{
			api_router:  NewRouter(),
			home_router: NewRouter(),
		},
	}
}

func (s *Server) Route(path string, handler http.HandlerFunc, router *Router) {
	router.rules[path] = handler
}

func (s *Server) AddMiddleware(fh http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	// add default middlewares
	middlewares = append(middlewares, LoggerMiddleware())
	for _, middleware := range middlewares {
		fh = middleware(fh)
	}
	return fh
}

func (s *Server) Listen() error {
	http.Handle("/api/", s.routers.api_router)
	http.Handle("/home/", s.routers.home_router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
