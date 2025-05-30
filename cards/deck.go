package main

import "fmt"

// Create a new type of deck which is a slice of string
type deck []string

// No receiver as you won't have an instance you are working with
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// Splitting a slice to make a hand from a deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// d -> Instance of the type deck that we are working with
// deck -> Type for which all variables of this type would be able to call this function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
