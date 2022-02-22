package main

import "fmt"

//Declarando o novo tipo: deck, que funciona como um slice de string.
type deck []string

//Lembrar que (d deck) é um receiver, que torna a função disponíveis para
//variáveis do tipo deck.
func (d deck) print() {
	for i, singleCard := range d {
		fmt.Println(i, singleCard)
	}
}

//Não tem receiver porque está justamente criando uma variável do tipo deck, ou seja
//não faria sentido que para criar um novo deck precisássemos de um já pronto.
//Recebe um parâmetro caso desejem que tenha coringas.
func newDeck(joker bool) deck {
	cards := deck{}

	cardSuits := []string{ "Spades", "Diamonds", "Hearts", "Clu"}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	//Iterando pelos nipes e valores para criar todas as combinações de cartas.
	for _, suit := range cardSuits{
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	if(joker){
		cards = append(cards, "Joker", "Joker")
	}

	return cards
}