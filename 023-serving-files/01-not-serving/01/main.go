package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(wr, `<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
}
