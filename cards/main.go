package main

func main() {
	cards := newDeck()
	cards.saveToFile("myDeck")
	// cards := newDeckFromFile("myDeck");
	cards.print()
	cards.shuffle()
	deal(cards, 5)
	cards.print()
}