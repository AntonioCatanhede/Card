package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {

		t.Errorf("Deck length must be 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {

		t.Errorf("First card must be Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {

		t.Errorf("Last card must be King of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {

		t.Errorf("Deck length must be 52, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
