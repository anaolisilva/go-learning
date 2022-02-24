package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck(false)

	if len(d) != 52 {
		t.Errorf("Expected deck lenght of 52, but got %v", len(d))
	}

	jokerDeck := newDeck(true)

	if len(jokerDeck) != 54 {
		t.Errorf("Expected joker deck lenght of 54, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of deck as 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of deck as 'King of Clubs', but got %v", d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	 os.Remove("_decktesting")

	 d := newDeck(false)

	 d.saveToFile("_decktesting")

	 loadedDeck := newDeckFromFile("_decktesting")

	 if len(loadedDeck) != 52 {
		 t.Errorf("Expected deck lenght of 52, but got %v", len(d))
	 }
}
