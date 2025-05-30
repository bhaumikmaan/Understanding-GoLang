package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck which is a slice of string
type deck []string

// No receiver as you won't have an instance you are working with
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
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

func (d deck) toString() string {
	//return strings.Join(d, ",")
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(filename string) error {
	// Pass in filename, convert deck to byte slice and pass the permission 0666 - anyone can read and write
	fmt.Println("Saving data to file:", filename)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// Log the error and exit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) shuffle() {
	// create a source for a new seed for randomizer
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Println("Shuffling...")
	for i := range d {
		// generate a random number from [0,i+1) i.e. 0 to ith number
		j := r.Intn(len(d))
		// swap ith element with jth element
		d[i], d[j] = d[j], d[i]
	}
}
