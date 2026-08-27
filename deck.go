package main

import "fmt"

//create a new type of 'deck'
//which is a slice of strings

type deck []string

func (d deck) print() { //any variable of type 'deck' now gets access to the "print" method
	for i, card:= range d {
		fmt.Println(i, card)
	}
}