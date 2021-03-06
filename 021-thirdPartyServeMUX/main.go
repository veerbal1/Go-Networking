package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyProcess)
	mux.GET("/user/:name", user)
	mux.GET("/blog/:category/:article", blogRead)
	mux.POST("/blog/:category/:article", blogWrite)
	http.ListenAndServe(":8080", mux)
}

func user(wr http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(wr, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(wr http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(wr, "READ CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(wr, "READ ARTICLE, %s!\n", ps.ByName("article"))
}

func blogWrite(wr http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(wr, "WRITE CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(wr, "WRITE ARTICLE, %s!\n", ps.ByName("article"))
}

func index(wr http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := tpl.ExecuteTemplate(wr, "index.gohtml", nil)
	HandleError(wr, err)
}

func about(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

// HandleError is for handling errors
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
