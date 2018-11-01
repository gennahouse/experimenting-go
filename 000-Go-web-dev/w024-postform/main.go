package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmp *template.Template

type hand int

func init() {
	tmp = template.Must(template.ParseFiles("tpl.gohtml"))
}

func (h hand) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tmp.ExecuteTemplate(w, "tpl.gohtml", req.PostForm)
}

func main() {
	var h hand
	http.ListenAndServe(":8080", h)
}
