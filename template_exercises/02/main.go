package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{
			"Marriot Marquis San Diego Marina",
			"333 West Harbor Drive",
			"San Diego", "92101", "Southern",
		},
		hotel{
			"The Ritz-Carlton, San Francisco",
			"600 Stockton Street",
			"San Francisco", "94108", "Central",
		},
		hotel{
			"Hampton Inn & Suites Redding",
			"2160 Larkspur Lane",
			"Redding", "96002", "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
