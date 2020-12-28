package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func dog(wr http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	fmt.Fprintf(wr, "Dog--")
}

func index(wr http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	fmt.Fprintf(wr, "Index")
}

func me(wr http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	fmt.Fprintf(wr, "Hello %s!\n", pr.ByName("name"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/dog", dog)
	mux.GET("/me/:name", me)

	http.ListenAndServe(":8080", mux)
}
