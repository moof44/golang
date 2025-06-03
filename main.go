package main

func main() {
	// cards, err := newDeckFromFile("my_cards")
	// if err != nil {
	// 	cards = newDeck()
	// }
	// cards.print()
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	cards := newDeck()
	cards.shuffle()
	cards.print()

}