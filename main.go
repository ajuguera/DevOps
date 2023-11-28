package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Soy alumno de la UOC - Método GET")
	case http.MethodPost:
		fmt.Fprintf(w, "Soy alumno de la UOC - Método POST")
	default:
		http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}