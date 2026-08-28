package main

func main() {
	cards := newDeck() //create a new deck
	cards.saveToFile("my_cards") //we trying to save the card deck into a filename called my_cards
}
