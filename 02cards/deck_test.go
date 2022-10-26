package main

import (
	"os"
	"testing"
)

func TestGetNewDeck(t *testing.T) {
	newDeck := getNewDeck()

	if len(newDeck) != 48 {
		t.Errorf("Expected deck to be of size 48, but got %v", len(newDeck))
	}

	if newDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first element of  deck to be Ace of Spades, but got %v", newDeck[0])
	}

	if newDeck[len(newDeck)-1] != "king of Diamonds" {
		t.Errorf("Expected last element of  deck to be king of Diamonds, but got %v", newDeck[0])
	}
}

func TestSaveToFileAndGetDeckFromFile(t *testing.T) {
	os.Remove("_deckTestingFile")

	newDeck := getNewDeck()

	newDeck.saveToFile("_deckTestingFile")

	deckFromDisk := getDeckFromFile("_deckTestingFile")

	if len(deckFromDisk) != 48 {
		t.Errorf("Error in Saving to file and reading from file.")
	}

	os.Remove("_deckTestingFile")

}
