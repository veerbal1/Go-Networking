package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func dog(wr http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	fmt.Fprintf(wr, "Dog--")
}

func index(wr http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	err := tpl.ExecuteTemplate(wr, "index.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func me(wr http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	err := tpl.ExecuteTemplate(wr, "me.gohtml", pr.ByName("name"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/dog", dog)
	mux.GET("/me/:name", me)

	http.ListenAndServe(":8080", mux)
}
