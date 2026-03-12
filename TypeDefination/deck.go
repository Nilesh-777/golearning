package main

import "fmt"

type deck []string

func newDeck() deck {
	cardSuits := deck{"Spade", "Diamond", "Heart", "Club"}
	cardNumbers := deck{"Ace", "One", "Two", "Three", "Four"}
	cards := deck{}

	for _, cardSuit := range cardSuits {
		for _, cardNumber := range cardNumbers {
			cards = append(cards, cardSuit+" of "+cardNumber)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
