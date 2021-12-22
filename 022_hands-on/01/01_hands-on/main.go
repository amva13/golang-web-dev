package main

import (
	"html/template"
	"io"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func cat(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("022_hands-on/01/03_hands-on/mytemplate.gohtml"))
	tpl.ExecuteTemplate(res, "mytemplate.gohtml", 42)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) { io.WriteString(w, "hi") })
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":8080", nil)
}
