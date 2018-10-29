package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

var fn = template.FuncMap{
	"x": mol,
	"d": div,
	"a": add,
}

func init() {
	tmp = template.Must(template.New("").Funcs(fn).ParseFiles("index.gohtml"))
}

func main() {
	err := tmp.ExecuteTemplate(os.Stdout, "index.gohtml", 20)
	if err != nil {
		log.Fatalln(err)
	}
}

func mol(i int) int {
	i *= 7
	return i
}

func div(i int) int {
	i /= 2
	return i
}

func add(i int) int {
	i += 92
	return i
}
