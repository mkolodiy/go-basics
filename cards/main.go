package main

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, newCard())

	cards.print()
	someOtherPrint(cards)
}

func newCard() string {
	return "Five of Diamonds"
}