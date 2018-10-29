package main

import (
	"os"
	"text/template"
)

var tmp *template.Template

type car struct {
	Model string
	Year  int
}

func init() {
	tmp = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	aventador := car{
		Model: "aventador",
		Year:  2018,
	}
	nine11 := car{
		Model: "911",
		Year:  2015,
	}

	cars := []car{
		aventador,
		nine11,
	}

	tmp.ExecuteTemplate(os.Stdout, "index.gohtml", cars)
}
