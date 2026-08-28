package main

import "fmt"

func main() {
	cards := newDeck() //create a new deck

	fmt.Println(cards.toString()) //convert the cards deck to slice of strings by calling the toString then display it
}
