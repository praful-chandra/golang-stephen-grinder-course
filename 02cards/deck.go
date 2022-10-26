package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func getNewDeck() deck {
	newDeck := deck{}

	cardValues := []string{"Ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "queen", "king"}
	cardTitles := []string{"Spades", "Hearts", "Clubs", "Diamonds"}

	for _, ct := range cardTitles {
		for _, cv := range cardValues {
			value := cv + " of " + ct
			newDeck = append(newDeck, value)
		}
	}

	return newDeck
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) {

	newFile, newFileErr := os.Create(fileName)

	if newFileErr != nil {
		fmt.Println("Error occured in creating new file : \n", newFileErr)
		os.Exit(1)
	}

	deckString := strings.Join(d, ",")

	byteDeck := []byte(deckString)

	saveFileLen, saveFileErr := newFile.Write(byteDeck)

	if saveFileErr != nil {
		fmt.Println("Error saving file to disk: \n", saveFileErr)
		os.Exit(2)
	}

	fmt.Println("File saved successfully : ", saveFileLen)
}

func getDeckFromFile(fileName string) deck {

	byteDeck, readFileError := os.ReadFile(fileName)

	if readFileError != nil {
		fmt.Println("Error reading file from disk : \n", readFileError)
		os.Exit(3)
	}

	return deck(strings.Split(string(byteDeck), ","))
}

func (d deck) shuffle() {

	newSource := rand.NewSource(time.Now().UnixNano())

	newRand := rand.New(newSource)

	for i := range d {

		rp := newRand.Int63n(int64(len(d)))

		d[i], d[rp] = d[rp], d[i]
	}
}

func (d deck) print() {
	for i, val := range d {
		fmt.Println(i, val)
	}
}
