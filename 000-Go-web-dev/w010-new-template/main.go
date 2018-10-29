package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmp := template.New("this")
	tmp, err := tmp.Parse("Hello from this")
	if err != nil {
		log.Fatalln(err)
	}
	tmp.ExecuteTemplate(os.Stdout, "this", nil)
}
