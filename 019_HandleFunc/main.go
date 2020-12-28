package main

import (
	"fmt"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am Dog")
}

func c(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am cat")
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "I am animal")
	})
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
