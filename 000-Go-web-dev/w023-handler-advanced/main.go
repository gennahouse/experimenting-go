package main

import (
	"log"
	"net/http"
	"text/template"
)

type hand int

var tmp *template.Template

func (h hand) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	var h hand
	http.ListenAndServe(":8020", h)
}
