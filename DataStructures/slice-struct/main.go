package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Budha",
		Motto: "The belief of no beliefs",
	}
	god := sage{
		Name:  "WaheGuru",
		Motto: "God is One",
	}
	sages := []sage{buddha, god}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
