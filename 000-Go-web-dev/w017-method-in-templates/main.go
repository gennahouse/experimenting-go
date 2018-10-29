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
	p := person{
		Name: "Alex",
		Age:  64,
	}
	tmp.Execute(os.Stdout, p)
}

func (p person) DecAge() int {
	return p.Age + 10
}

func (p person) Ten(i int) int {
	return i - p.Age
}
