package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Budha",
		"Love":     "Jesus",
		"Prophet":  "Muhamad"}

	// nf, eror := os.Create("index.html")
	// if eror != nil {
	// 	log.Fatalln(eror)
	// }
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
