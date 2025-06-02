package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{lastName: "Alex", firstName: "Anderson"}
	var emptyStruct person
	// Print Struct
	fmt.Println(alex)
	// Print all the fields and values of the struct
	fmt.Printf("%+v\n", alex)
	// Print a variable struct with 0 values
	fmt.Println(emptyStruct)
	emptyStruct.firstName = "Hi"
	emptyStruct.lastName = "Empty"
	fmt.Println(emptyStruct)

	// Since we added a new struct of contact in person - all the above structs will have 0 value for contact info
	jim := person{firstName: "Jim", lastName: "Anderson", contactInfo: contactInfo{
		email: "jim@gmail.com",
		zip:   42,
	}}
	jim.print()
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
