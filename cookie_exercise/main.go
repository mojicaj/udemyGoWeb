package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		})
	} else {
		v, _ := strconv.Atoi(c.Value)
		http.SetCookie(w, &http.Cookie{
			Name:  "my-cookie",
			Value: strconv.Itoa(v + 1),
		})
	}
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}
