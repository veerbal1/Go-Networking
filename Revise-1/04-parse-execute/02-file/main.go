package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}

	defer nf.Close()
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal(err)
	}
}
