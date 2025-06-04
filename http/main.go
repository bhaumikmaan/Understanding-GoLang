package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// https://pkg.go.dev/net/http#Get
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Found Error: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	// https://pkg.go.dev/net/http#Response
	fmt.Println(resp)

	// Read all data into a byte slice once
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		os.Exit(1)
	}
	// 1. Returns an HTML String output from google.com
	fmt.Println(string(bodyBytes))

	fmt.Println("------------------------- Reader Writer Interface ---------------------------------------")
	// TODO: Pass the response body to a reader and get a byte slice of data & print

	// 2. Use io.Copy to write to os.Stdout
	io.Copy(os.Stdout, bytes.NewReader(bodyBytes))

	fmt.Println("------------------------- Custom Writer ---------------------------------------")
	// TODO: Implement our own Write function to create a customer writer to be used in io.copy
	// 3. Use a custom writer
	lw := logWriter{}
	io.Copy(lw, bytes.NewReader(bodyBytes))
}

func (logWriter) Write(bs []byte) (int, error) {
	// Incorrect random return values that makes the code compile but doesn't make sense in terms of actual implementation
	// return 1, nil
	fmt.Println(string(bs))
	fmt.Println("Printed out: ", len(bs), "bytes")
	return len(bs), nil
}
