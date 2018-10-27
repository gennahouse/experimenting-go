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
	names := []string{
		"Robert",
		"Carl",
		"Jamal",
	}

	tmp.ExecuteTemplate(os.Stdout, "index.gohtml", names)
}
