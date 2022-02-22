package main

import "fmt"

func main() {

	//Declarando e populando um novo tipo deck.
	cards := newDeck(false)

	hand, remainingCards := deal(cards, 3)

	fmt.Println("Hand dealt:")
	//Chamando a função criada no tipo deck
	hand.print()
	
	fmt.Println("Remaining cards:")
	remainingCards.print()

}
