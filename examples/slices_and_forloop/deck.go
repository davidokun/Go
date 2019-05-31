package main

import "fmt"

type deck []string

// Receiver for type deck
// Convention en Go to use only 1 letter for the receiver parameter
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
