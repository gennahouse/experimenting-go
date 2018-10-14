package main

import "fmt"

func main() {
	s := "favSport"
	switch s {
	case "no":
		fmt.Println("Option one")
	case "footbal":
		fmt.Println("Option two")
	case "favSport":
		fmt.Println("Winner")
	case "imfat":
		fmt.Println("You need to join a team")
	}
}
