package main

import (
	"fmt"
	"net/http"
)

type gool int

func (g gool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey")
}

func main() {
	var gf gool
	http.ListenAndServe(":8000", gf)
}
