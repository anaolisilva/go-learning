package main

import "fmt"

type deck []string

func newDeck(joker bool) deck {
	cards := deck{}

	cardSuits := []string{ "Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

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

func (d deck) print() {
	for i, singleCard := range d {
		fmt.Println(i, singleCard)
	}
}

func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}