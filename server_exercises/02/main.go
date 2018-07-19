package main

import (
	"io"
	"net/http"
)

func h(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "You're Home!")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func m(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, req.URL.Path[4:])
}

func main() {
	http.HandleFunc("/", h)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)

	http.ListenAndServe(":8080", nil)
}
