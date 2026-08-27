package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()} //the deck is from the deck.go //it is slice //remeber when we run we have to include deck.go, that is go run main.go deck.go
	cards = append(cards, "Six of Spades")      //to add additional elements to it //it returns new slice so we can store it on the variable and it does not destroy the old one

	// fmt.Println(cards)

	cards.print() //cards is passed as d in the deck.go //as a copy in short

}

func newCard() string {
	return "Five of Diamonds"
}
