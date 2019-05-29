package main

import "fmt"

func main() {
	// Verbose variable initialization
	var card string = "Ace of Spades"
	fmt.Println(card)
	card = "Daimonds"

	// Simplified variable initialization
	// Go "infer" the type base on the assigment
	otherCard := "Hearths"
	fmt.Println(otherCard)

	// Can't assign a different type
	//otherCard = 1
}
