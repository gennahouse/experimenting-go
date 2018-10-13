/*This will be based on []string type
I will re-do that with structs as well
That way  I can get the user to get more cards and play blackjack
*/
package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.addCards()
	fmt.Println(cards)
	cardsUser, cardsDealer := cards.newHand()
	fmt.Println("User has:", cardsUser)
	fmt.Println("Dealer has;", cardsDealer)
}
