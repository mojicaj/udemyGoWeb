package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func h(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func d(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func m(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "name.gohtml", req.URL.Path[4:])
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(h))
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/me/", http.HandlerFunc(m))

	http.ListenAndServe(":8080", nil)
}
