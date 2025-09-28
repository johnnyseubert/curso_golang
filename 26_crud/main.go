package main

import (
	"crud/servidor"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

	fmt.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
