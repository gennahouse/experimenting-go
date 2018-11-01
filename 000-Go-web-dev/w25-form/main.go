package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmp *template.Template

type winston int

func (p winston) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tmp.Execute(w, req.Form)
}
func init() {
	tmp = template.Must(template.ParseFiles("form.gohtml"))
}

func main() {
	var w winston

	http.ListenAndServe(":8081", w)
}
