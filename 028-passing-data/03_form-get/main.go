package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		v := r.FormValue("q")
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(rw, `
		<form method="get">
		<input type="text" name="q"/>
		<input type="submit" />
		</form>`+v)
	})
	http.ListenAndServe(":8080", nil)
}
