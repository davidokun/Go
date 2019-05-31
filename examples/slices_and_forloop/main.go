package main

import "fmt"

func main()  {
	cards := []string{"Ace of Spades", getCard(), "Jack of Souls"}
	cards = append(cards, "3 of Hearts")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func getCard() string {
	return "Five of Diamonds"
}
