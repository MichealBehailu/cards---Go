package main

import "fmt"

//create a new type of 'deck'
//which is a slice of strings

type deck []string

func newDeck() deck { //return a value of deck
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} //these are the variants
	cardValues := []string{"Ace", "Two", "Three", "Four"}          //these are the values from ace to king

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { //any variable of type 'deck' now gets access to the "print" method
	for i, card := range d {
		fmt.Println(i, card)
	}
}
