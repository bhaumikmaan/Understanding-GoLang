package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the file that whose name is given by os.Args
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error Found:", err)
		os.Exit(1)
	}
	// Read and stdOut the data in file f
	io.Copy(os.Stdout, f)
}
