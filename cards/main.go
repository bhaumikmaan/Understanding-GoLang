package main

import (
	"fmt"
)

// deckSize := 20 -> This is incorrect as a variable CAN be defined outside a function but CAN NOT be assigned a value

func main() {
	// var card string = "Ace of Spades". Can also be defined as follows
	card := "Ace of Spades"
	fmt.Println(card)

	// Reassign without using the := operator
	card = reassignCard()
	fmt.Println(card)
}

// if the type 'string' is not mentioned, this will throw an error
func reassignCard() string {
	return "Five of Diamonds"
}
