package main

func main() {
	cards := newDeck()

	hand, remaningCards := deal(cards, 5) //two separate variable needed two store the two returned value

	hand.print() //we can call the print function because the two varaibles are type of deck
	remaningCards.print()
}
