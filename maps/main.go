package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	delete(colors, "red")
	fmt.Println(colors)
	printMap(colors)

	// Zero value map - No key value pairs, can be added on the go
	// Method 1
	// var emptyMap map[string]string
	// Method 2
	emptyMap := make(map[string]string)
	fmt.Println(emptyMap)
}

// Reference type so no need to handle pointers
func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Color Found: ", key, " with hex value: ", value)
	}
}
