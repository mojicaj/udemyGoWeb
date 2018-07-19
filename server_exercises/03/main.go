package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func h(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func d(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func m(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "name.gohtml", req.URL.Path[4:])
}

func main() {
	http.HandleFunc("/", h)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)

	http.ListenAndServe(":8080", nil)
}
