package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmp *template.Template

type hand int

var fc = template.FuncMap{
	"test": lol,
}

func lol() int {
	log.Println("hello")
	return 10
}

func init() {
	tmp = template.Must(template.New("").Funcs(fc).ParseFiles("test.gohtml"))
}

func (h hand) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln("there was a problem with the parsing")
	}
	tmp.ExecuteTemplate(w, "test.gohtml", req.PostForm)
}

func main() {
	var k hand
	http.ListenAndServe(":8090", k)
}
