package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

func main() {

	kl := template.Must(template.ParseGlob("*.html"))
	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Página Raiz"))
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar usuários"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}
