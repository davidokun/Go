package main

func main()  {
	//Use the declared type in "deck.go"
	cards := deck{"Ace of Spades", getCard(), "Jack of Souls"}
	cards = append(cards, "3 of Hearts")

	//Use the defined receiver for type deck
	cards.print()
}

func getCard() string {
	return "Five of Diamonds"
}
