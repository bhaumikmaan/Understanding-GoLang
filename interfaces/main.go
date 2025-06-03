package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	// part of the interface, now type spanish and english bot both have this function
	// thus becoming a part of this interface proprietorially
	// Once that happens they get access to functions like printGreeting that require a "bot" type to be passed
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// NOTE: If we make same function-same functionality for each different bot - it is not ideal hence the concept of interfaces
//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())
//}
//
//// GO doesn't allow method overloading like Java
//func printGreeting(sb spanishBot) {
//	fmt.Println(sb.getGreeting())
//}

func (englishBot) getGreeting() string {
	return "Hello"
}
func (spanishBot) getGreeting() string {
	return "Hola"
}
