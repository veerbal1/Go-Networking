package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		v := r.FormValue("q")
		io.WriteString(rw, "Do I search:"+v)
	})
	http.ListenAndServe(":8080", nil)
}
