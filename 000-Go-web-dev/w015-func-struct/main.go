package main

import (
	"os"
	"text/template"
)

var tmp *template.Template

var fn = template.FuncMap{
	"fullname": fullName,
}

type person struct {
	name    string
	surname string
}

func init() {
	tmp = template.Must(template.New("").Funcs(fn).ParseFiles("index.gohtml"))
}

func main() {
	george := person{
		name:    "George",
		surname: "TheSecond",
	}
	tmp.ExecuteTemplate(os.Stdout, "index.gohtml", george)
}

func fullName(p person) string {
	name := p.name + " " + p.surname
	return name
}
