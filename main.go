package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" //one way to define

	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
