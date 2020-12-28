package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	var str string

	fmt.Println(r.Method)

	if r.Method == http.MethodPost {
		// open
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer f.Close()

		// For your information
		fmt.Println("\nFile:", f, "\nHeader", h, "\nError", err)

		// Read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}

		str = string(bs)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" action="/" enctype="multipart/form-data">
		<input name="q" type="file"/>
		<input type="submit"/>
	</form><br>
	`+str)

}
