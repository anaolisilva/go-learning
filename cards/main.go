package main

import "fmt"

func main() {

	//Declarando e populando um slice.
	cards := []string{ newCard(), "Eight of Hearts"}
	
	//Adicionando elemento ao slice (não altera, constrói um novo).
	cards = append(cards, "Six of Spades")

	//Iterando pelo slice de cards (lembrar de usar o index, se não o código não compila)
	for i, singleCard := range cards {
		fmt.Println(i, singleCard)
	}
}

func newCard() string {
	return "Five of Diamonds"
}