package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var examples []exCompose

type exStruct struct {
	name  string
	power int
}

type exCompose struct {
	*exStruct // composition
	isGood    bool
}

func exFunc1(ex exCompose) string {
	return ex.name
}

func exFunc2(ex exCompose) int {
	return ex.power
}

func init() {
	fMap := template.FuncMap{
		"getName":   exFunc1,
		"getPower":  exFunc2,
		"getStruct": func(ex exCompose) exCompose { return ex },
	}
	tpl = template.Must(template.New("").Funcs(fMap).ParseGlob("*.gohtml"))
	exA := exStruct{"A", 5}
	exB := exStruct{"B", 6}
	exC := exStruct{"C", 7}
	examples = []exCompose{{&exA, true}, {&exB, false}, {&exC, true}}
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", examples)
	if err != nil {
		log.Fatalln(err)
	}
}
