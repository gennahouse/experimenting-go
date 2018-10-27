package main

import (
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	name := "Robert"
	tmp.ExecuteTemplate(os.Stdout, "index.gohtml", name)
}
