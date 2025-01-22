package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	templates = template.Must(template.ParseGlob("*.html"))

	u := usuario{"Ricardo", "kreycek@gmail.com"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", u)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}
