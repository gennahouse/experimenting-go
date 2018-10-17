package main

import "fmt"

func main() {
	mds := make([][]string, 0, 10)

	ss := make([]string, 0, 10)

	for i := 0; i < 9; i++ {
		ss = append(ss, "X")
		mds = append(mds, ss)
	}

	for _, v := range mds {
		fmt.Println(v)
	}
}
