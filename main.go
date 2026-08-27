package main

import (
	"fmt"
)

func main() {
	cards := []string{"Ace of Diammonds", newCard()} //inside the brace we can put elements to it
	cards = append(cards, "Six of Spades")                      //to add additional elements to it //it returns new slice so we can store it on the variable and it does not destroy the old one

	// fmt.Println(cards)

	//to iterate over the slice
	for i, card := range cards { //range keyword is used to iterate over every element in the slice
		fmt.Println(i, card)
	}
	
}

func newCard() string {
	return "Five of Diamonds"
}
