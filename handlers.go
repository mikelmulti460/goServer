package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola mundo 2")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "api :3")
}

func AuthenticationError(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Error! no estás autenticado")
}

func PageNotFoundError(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Se nos perdió esta página :c")
}
