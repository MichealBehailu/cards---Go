package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

//create a new type of 'deck'
//which is a slice of strings

type deck []string

func newDeck() deck { //return a value of deck
	cards := deck{} //this means cards is the type of deck

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} //these are the variants
	cardValues := []string{"Ace", "Two", "Three", "Four"}          //these are the values from ace to king

	for _, suit := range cardSuits { //the idea of this loop is to add both the cardSuits and also the cardValues to cards
		for _, value := range cardValues { //the _ is the replace for variable that we dont use
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

func deal(d deck, handSize int) (deck, deck) { //the (deck, deck) means go expect two things to be returned
	return d[:handSize], d[handSize:] //for the first return value (d[:handSize]) it means from starting(0) to the handover //the second return value (d[handSize:]) is from handsize to the end
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") //the []string(d) - this will convert the deck to slice of strings (or in one string) //and Join is used to join the slice of strings //and the "," - is the separator //refer package on go documentation(pkg)
}

func (d deck) saveToFile(filename string) error { //it takes a filename and also a reciver deck
	return os.WriteFile(filename, []byte(d.toString()), 0666) // 0666 - means anyone can read and write this file, it is like permission
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) //bs inthis case it is byteslice

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) //if 0 is passed indicates success //but non zero means an error occured and the program terminates immediately
	}

	s := strings.Split(string(bs), ",") 
	return deck(s)

}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1) //intn - is used to generate random number 
		
		d[i], d[newPosition] = d[newPosition], d[i] //swap
	}
}