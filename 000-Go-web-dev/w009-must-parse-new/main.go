package main

import (
	"os"
	"text/template"
)

func main() {
	tmp := template.Must(template.New("test").Parse("Hello from test"))
	tmp.ExecuteTemplate(os.Stdout, "test", nil)
}
