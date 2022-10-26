package main

import "fmt"

func main() {
	fmt.Println("Card App..")

	// cards := getNewDeck()

	existingDeck := getDeckFromFile("my_cards")
	existingDeck.shuffle()
	existingDeck.print()

}
