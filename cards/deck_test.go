package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	//Primeiros dois conferem tamanho do deck
	d := newDeck(false)

	if len(d) != 52 {
		t.Errorf("Expected deck lenght of 52, but got %v", len(d))
	}

	jokerDeck := newDeck(true)

	if len(jokerDeck) != 54 {
		t.Errorf("Expected joker deck lenght of 54, but got %v", len(d))
	}

	//Próximos dois conferem integridade das cartas
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of deck as 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of deck as 'King of Clubs', but got %v", d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	
	//Apaga quaisquer arquivos com esse nome em caso de falha no meio do teste.
	//Go não realiza limpeza dos arquivos de teste.
	 os.Remove("_decktesting")

	 d := newDeck(false)

	 d.saveToFile("_decktesting")

	 loadedDeck := newDeckFromFile("_decktesting")

	 //Confere tamanho do deck carregado a partir de arquivo criado.
	 if len(loadedDeck) != 52 {
		 t.Errorf("Expected deck lenght of 52, but got %v", len(d))
	 }
}
