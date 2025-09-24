package main

import (
	"log"
	"net/http"
)

func renderHtml(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "arquivos/index.html")
}

func main() {
	http.HandleFunc("/", renderHtml)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
