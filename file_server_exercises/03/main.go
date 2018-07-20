package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", h)
	http.Handle("/pics/", fs)

	http.ListenAndServe(":8080", nil)
}

func h(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
