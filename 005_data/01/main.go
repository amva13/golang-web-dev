package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var examples []string

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	examples = []string{"donuts", "fire", "golang"}
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", examples)
	if err != nil {
		log.Fatalln(err)
	}
}
