package main

import "fmt"

func main() {

	cards := newDeckFromFile("my_cards")

	hand, remainingCards := deal(cards, 3)

	fmt.Println("Hand dealt:")
	hand.print()
	
	fmt.Println("Remaining cards:")
	remainingCards.print()

	fmt.Print(cards.toString())
}
