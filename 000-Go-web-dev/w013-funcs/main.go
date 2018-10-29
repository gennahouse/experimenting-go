package main

import (
	"os"
	"text/template"
)

var tmp *template.Template

var fn = template.FuncMap{
	"x": molt,
}

func init() {
	tmp = template.Must(template.New("").Funcs(fn).ParseFiles("index.gohtml"))
}

func main() {
	tmp.ExecuteTemplate(os.Stdout, "index.gohtml", 10)
}

func molt(i int) int {
	i *= 10
	return i
}
