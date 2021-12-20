package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var examples []exStruct

type exStruct struct {
	name  string
	power int
}

func exFunc1(ex exStruct) string {
	return ex.name
}

func exFunc2(ex exStruct) int {
	return ex.power
}

func init() {
	fMap := template.FuncMap{
		"getName":  exFunc1,
		"getPower": exFunc2,
	}
	tpl = template.Must(template.New("").Funcs(fMap).ParseFiles("tpl.gohtml"))

	examples = []exStruct{{"A", 5}, {"B", 6}, {"C", 7}}
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", examples)
	if err != nil {
		log.Fatalln(err)
	}
}
