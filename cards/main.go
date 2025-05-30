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
	//--------------------------------------------------------------------------------------------------------------------//
	fmt.Println("------------------------- FILE MANAGEMENT & TYPE CONVERSIONS ---------------------------------------")
	greeting := "Hi There! This is a string"
	fmt.Println([]byte(greeting))

	// SAVING DECK TO A FILE
	// Deck -> []string -> string -> []byte
	fmt.Println(cardsType.toString())
	hand.saveToFile("hand.txt")

	// READING FROM A FILE
	// read bytes and convert it back to deck
	retrievedHand := newDeckFromFile("hand.txt")
	retrievedHand.print()
	errorHand := newDeckFromFile("randomFile.txt")
	errorHand.print()
}

// if the type 'string' is not mentioned, this will throw an error
func reassignCard() string {
	return "Five of Diamonds"
}
