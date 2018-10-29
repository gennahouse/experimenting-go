package main

import (
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	tmp.ExecuteTemplate(os.Stdout, "index.gohtml", nil)
}
