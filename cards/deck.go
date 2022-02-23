package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toString() string {	
	stringSlice := []string(d)

	return strings.Join(stringSlice, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	deckByteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("Quitting program.")
		os.Exit(1)
	}

	s := strings.Split(string(deckByteSlice), ",")
	deckFromFile := deck(s)

	return deckFromFile
}