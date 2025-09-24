package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Home Page!"))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Users Page!"))
}

func main() {
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/usuarios", usersHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
