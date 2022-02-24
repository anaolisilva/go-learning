package main

import "fmt"

func main() {

	//Declarando e populando um novo tipo deck.
	//cards := newDeck(false)

	//Salva cartas num arquivo.
	//cards.saveToFile("my_cards")

	//Cria novo deck a partir de arquivo
	cards := newDeckFromFile("my_cards")
	
	cards.shuffle()

	hand, remainingCards := deal(cards, 3)

	fmt.Println("Hand dealt:")
	//Chamando a função criada no tipo deck
	hand.print()
	
	fmt.Println("Remaining cards:")
	remainingCards.print()
}