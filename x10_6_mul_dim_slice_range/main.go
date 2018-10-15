package main

import (
	"fmt"
)

func main() {
	jb := []string{
		"James",
		"Bond",
		"Shaken not stirred",
	}
	mp := []string{
		"Miss",
		"Moneypenny",
		"Helloooooo James.",
	}
	mds := [][]string{jb, mp}
	fmt.Println(mds)
}
