package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Bao Bao Bao")
}

type hotcat int

func (m hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Meaoo Meaoo")
}
func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", d)
}
