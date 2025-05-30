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
	//--------------------------------------------------------------------------------------------------------------------//
	fmt.Println("----------------------------- LEARNING ABOUT SLICES -------------------------------------------")

	// Arrays and Slice
	cards := []string{"Ace of Spades", reassignCard()}
	fmt.Println(cards)

	cards = append(cards, "Ace of Clubs") // Note -> Append makes a new slice and is then reassigned to cards
	for index, card := range cards {
		fmt.Println(index, card)
	}
	//--------------------------------------------------------------------------------------------------------------------//
	fmt.Println("----------------------------- LEARNING ABOUT TYPES -------------------------------------------")
	cardsType := newDeck()
	cardsType.print()
	hand, remainingHand := deal(cardsType, 3)
	hand.print()
	remainingHand.print()
}

// if the type 'string' is not mentioned, this will throw an error
func reassignCard() string {
	return "Five of Diamonds"
}
