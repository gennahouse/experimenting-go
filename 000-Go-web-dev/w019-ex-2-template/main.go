package main

import (
	"os"
	"text/template"
)

var tmp *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

func init() {
	tmp = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h1 := hotel{
		Name:    "Palace",
		Address: "Main Street 89",
		City:    "California",
		Zip:     80782,
		Region:  "Central",
	}
	h2 := hotel{
		Name:    "Imperial",
		Address: "Grove St 2839",
		City:    "California",
		Zip:     80732,
		Region:  "Southern",
	}
	h3 := hotel{
		Name:    "Velux",
		Address: "Fitfh evenue 621",
		City:    "California",
		Zip:     80712,
		Region:  "Northern",
	}
	hotels := []hotel{
		h1,
		h2,
		h3,
	}
	tmp.Execute(os.Stdout, hotels)
}
