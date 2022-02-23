package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Declarando o novo tipo: deck, que funciona como um slice de string.
type deck []string


//Não tem receiver porque está justamente criando uma variável do tipo deck, ou seja
//não faria sentido que para criar um novo deck precisássemos de um já pronto.
//Recebe um parâmetro caso desejem que tenha coringas.
func newDeck(joker bool) deck {
	cards := deck{}

	cardSuits := []string{ "Spades", "Diamonds", "Hearts", "Clubs"}

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


//Lembrar que (d deck) é um receiver, que torna a função disponível para
//variáveis do tipo deck.
func (d deck) print() {
	for i, singleCard := range d {
		fmt.Println(i, singleCard)
	}
}

//Função que retorna dois recursos diferentes, ambos do tipo deck.
func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Função para transformar o deck em umaString única, para que possamos transformá-lo em
//byteSlice e usar na função da biblioteca padrão de go writeFile.
func (d deck) toString() string {
	
	//transforma deck em []string (essencialmente seu tipo nativo)
	stringSlice := []string(d)

	//Função de uma biblioteca nativa de Go
	return strings.Join(stringSlice, ",")
}

//Função que vai salvar o deck de cartas num arquivo
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//Função que retorna um deck a partir de um arquivo.
func newDeckFromFile(filename string) deck {
	//Duas variáveis: uma guarda o byteSlice a ser recebido, a outra guarda o erro *se houver*.
	deckByteSlice, err := ioutil.ReadFile(filename)

	//nil é o tipo nulo de Go.
	if err != nil {
		//Loga erro e sai do programa.
		fmt.Println("Error: ", err)
		fmt.Println("Quitting program.")
		os.Exit(1)
		//Mais uma função de biblioteca padrão.
	}

	//Caminho oposto da função saveToFile. Transforma byteSlice em string única, depois as separa.
	s := strings.Split(string(deckByteSlice), ",")
	deckFromFile := deck(s)

	return deckFromFile
}