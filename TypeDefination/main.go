package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("Inital Deck")

	// Print all cards
	cards.print()

	// Split & Print
	hand, remainingCards := deal(cards, 2)
	fmt.Println("Deck1: ")
	hand.print()
	fmt.Println("Deck2: ")
	remainingCards.print()

	fmt.Print(cards.toString())

	cards.saveToFile("cards_content.txt")
}
