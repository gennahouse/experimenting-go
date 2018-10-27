package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	addison := person{
		Name: "Addison",
		Age:  26,
	}
	err := tmp.ExecuteTemplate(os.Stdout, "index.gohtml", addison)
	if err != nil {
		log.Fatalln(err)
	}
}
