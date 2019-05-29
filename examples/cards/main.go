package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)

	sum := sumNumbers(2, 3)
	fmt.Println(sum)
}

func sumNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func newCard() string {
	return "Ace of Spades"
}
