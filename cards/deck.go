package main

import "fmt"

// Create a new type of deck which is a slice of string
type deck []string

// d -> Instance of the type deck that we are working with
// deck -> Type for which all variables of this type would be able to call this function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
