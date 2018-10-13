package main

import (
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	d := deck{}
	return d
}

func (d *deck) addCards() {
	var x []string
	cardsValues := []string{
		"ace",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"jack",
		"queen",
		"king",
	}
	cardsTypes := []string{
		"spades",
		"diamonds",
		"hearts",
		"clubs",
	}
	for _, value := range cardsValues {
		for _, types := range cardsTypes {
			x = append(x, value+" of "+types)
		}
	}
	*d = x
}

func (d deck) newHand() ([]string, []string) {
	var cU []string
	var cD []string
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	for i := 0; i < 2; i++ {
		x := r.Intn(len(d) - 1)
		cU = append(cU, d[x])
	}
	for j := 0; j < 2; j++ {
		x := r.Intn(len(d) - 1)
		cD = append(cD, d[x])
	}
	return cU, cD
}
