package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

var fm = template.FuncMap{
	"fDateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
