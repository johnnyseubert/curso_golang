package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type User struct {
	Name  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		user := User{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}

		templates.ExecuteTemplate(writer, "index.html", user)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
