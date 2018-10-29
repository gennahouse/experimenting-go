package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tmp *template.Template

type timeToday struct {
	Today   string
	Timenow string
}

func init() {
	tmp = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	nowt := timeToday{
		Today:   time.Now().Format("02-01-2006"),
		Timenow: time.Now().Format("03:04:05"),
	}
	err := tmp.ExecuteTemplate(os.Stdout, "index.gohtml", nowt)
	if err != nil {
		log.Fatalln(err)
	}
}
