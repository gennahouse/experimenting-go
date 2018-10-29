package main

import (
	"os"
	"text/template"
)

var tmp *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	tmp = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	p1 := person{
		Name: "George",
	}

	p2 := person{
		Name: "Winston",
		Age:  21,
	}

	p3 := person{
		Age: 38,
	}

	people := []person{
		p1,
		p2,
		p3,
	}

	tmp.Execute(os.Stdout, people)
}
