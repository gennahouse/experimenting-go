package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmp *template.Template

type holo int

func (h holo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tmp.ExecuteTemplate(w, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func init() {
	tmp = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	var h holo
	err := http.ListenAndServe(":8070", h)
	if err != nil {
		log.Fatalln(err)
	}

}
