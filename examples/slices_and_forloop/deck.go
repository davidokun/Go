package main

import "fmt"

type deck []string

func newDeck() deck  {
	cards := deck{}

	cardsSuits := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	suitValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	// The "_" is used to ignore the use of the variable (compiler happy?)
	for _, suit := range cardsSuits {
		for _, value := range suitValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// Receiver for type deck
// Convention en Go to use only 1 letter for the receiver parameter
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
