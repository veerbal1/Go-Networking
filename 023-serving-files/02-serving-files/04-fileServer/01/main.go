package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(wr, `<!--not serving-->
	<img src="dog.jpg">`)
}
