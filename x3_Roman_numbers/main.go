package main

import "fmt"

func main() {
	var test int
	fmt.Println("Enter the number to convert:")
	fmt.Scanln(&test)
	u, d, c, m := decValues(test)
	fmt.Println("units", u)
	fmt.Println("tens", d)
	fmt.Println("houndreds", c)
	fmt.Println("thausand", m)
	num := convertRomans(u, d, c, m)
	fmt.Println("Roman numbers:", num)
}

func decValues(n int) (int, int, int, int) {
	units := 0
	tens := 0
	houndreds := 0
	thousand := 0
	for i := 0; i < n; i++ {
		units++
		//I can't use else if here
		if units >= 10 && tens < 10 {
			units = 0
			tens++
		}
		if tens >= 10 {
			units = 0
			tens = 0
			houndreds++
		}
		if houndreds >= 10 {
			units = 0
			tens = 0
			houndreds = 0
			thousand++
		}
	}
	return units, tens, houndreds, thousand
}

func convertRomans(u int, d int, c int, m int) string {
	units := map[int]string{
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		8: "VIII",
		9: "IX",
	}
	tens := map[int]string{
		1: "X",
		2: "XX",
		3: "XXX",
		4: "XL",
		5: "L",
		6: "LX",
		7: "LXX",
		8: "LXXX",
		9: "XC",
	}
	houndreds := map[int]string{
		1: "C",
		2: "CC",
		3: "CCC",
		4: "CD",
		5: "D",
		6: "DC",
		7: "DCC",
		8: "DCCC",
		9: "CM",
	}
	//Notice: I couldn't find the right characted for 5k and 10k so I used V and X
	thousand := map[int]string{
		1: "M",
		2: "MM",
		3: "MMM",
		4: "VM",
		5: "V",
		6: "VM",
		7: "VMM",
		8: "MVVV",
		9: "XM",
	}
	un := units[u]
	du := tens[d]
	cu := houndreds[c]
	mu := thousand[m]
	num := mu + cu + du + un
	return num
}
