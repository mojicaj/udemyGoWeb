package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Price string
}

type section struct {
	Section string
	Items   []item
}

type menu []section

type restaurant struct {
	Name, City string
	Menu       menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := restaurants{
		restaurant{
			"Metro Diner",
			"Jacksonville",
			menu{
				section{
					"Breakfast",
					[]item{
						item{"Omelette", "$5.99"},
						item{"Bagel", "$2.99"},
					},
				},
				section{
					"Lunch",
					[]item{
						item{"Cheeseburger", "$7.99"},
						item{"Salad", "$6.99"},
					},
				},
				section{
					"Dinner",
					[]item{
						item{"Steak & Veggies", "$15.99"},
						item{"Mac n Cheese", "$11.99"},
					},
				},
			},
		},
		restaurant{
			"Don Corleone's",
			"Savannah",
			menu{
				section{
					"Lunch",
					[]item{
						item{"Stromboli", "$7.99"},
						item{"Anti Pasta", "$6.99"},
					},
				},
				section{
					"Dinner",
					[]item{
						item{"Chicken Alfredo", "$12.99"},
						item{"Lasagna", "$13.99"},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
