package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmp, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tmp.Execute(os.Stdout, nil)
}
