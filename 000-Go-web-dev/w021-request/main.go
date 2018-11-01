package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmp *template.Template

type foo int

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tmp.ExecuteTemplate(w, "tpl.gohtml", r.Form)
}
func init() {
	tmp = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	var f foo
	http.ListenAndServe(":8000", f)
}
