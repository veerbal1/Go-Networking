package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "God is one")
	if err != nil {
		log.Fatal(err)
	}
}
