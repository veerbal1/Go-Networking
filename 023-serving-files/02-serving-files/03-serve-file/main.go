package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(wr, `<!--not serving-->
	<img src="/dog.jpg">`)
}

func dogPic(wr http.ResponseWriter, r *http.Request) {
	http.ServeFile(wr, r, "dog.jpg")
}
