package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func dog(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(wr, "Dog--")
}

func index(wr http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(wr, "index.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func me(wr http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(wr, "me.gohtml", "Mandeep")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
