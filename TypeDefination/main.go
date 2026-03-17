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

	fmt.Println()

	fmt.Println("Comma Separated Deck: ")
	fmt.Println(cards.toString())

	fmt.Println()

	fmt.Println("Deck from file: ")
	fileDeck := cards.readFromFile("cards_content.txt")
	fmt.Println(fileDeck)

	fmt.Println()

	fmt.Println("Shuffled Deck")
	fileDeck.shuffle()
	fmt.Println(fileDeck)
}
