package main

import "testing"

func TestNewDeck(t *testing.T) { //t - is the test handler
	d := newDeck()

	//the number of cards inside should be 16 if not
	if len(d) != 16 {
		len := len(d)
		t.Errorf("Expected deck length of 16, but got %v", len) //%v helps us to integrate the argument (len) with the string
	}
}
