package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Price string
}

type menu struct {
	Section string
	Items   []item
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := []menu{
		menu{
			"Breakfast",
			[]item{
				item{"Omelette", "$5.99"},
				item{"Bagel", "$2.99"},
			},
		},
		menu{
			"Lunch",
			[]item{
				item{"Cheeseburger", "$7.99"},
				item{"Salad", "$6.99"},
			},
		},
		menu{
			"Dinner",
			[]item{
				item{"Steak & Veggies", "$15.99"},
				item{"Mac n Cheese", "$11.99"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
