package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.bing.com",
		"http://www.golang.org",
		"http://www.stackoverflow.com",
	}

	ch := make(chan string)

	fmt.Println("------------------------- Status Check once ---------------------------------------")
	for _, link := range links {
		//fmt.Println(link)
		go checkLink(link, ch)
	}

	// Wait for all the links to complete and listen to channel until then
	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch + " Status Checked")
	}

	fmt.Println("------------------------- Status Check Repetition ---------------------------------------")
	for _, link := range links {
		//fmt.Println(link)
		go checkLink(link, ch)
	}
	// Listens to the channel indefinitely
	for l := range ch {
		go func(l string) {
			fmt.Println("Sleeping for 5 seconds for the link ", l, "...")
			time.Sleep(5 * time.Second)
			go checkLink(l, ch)
		}(l) // By doing this, we pass link as an argument making it as a copy from the main routine
	}
}

func checkLink(link string, ch chan string) {
	// Wait for 5 seconds for each fetch of the URL and not stop the MAIN routine
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "domain might be down:", err)
		ch <- link
		return
	}
	fmt.Println(link, " domain is up")
	ch <- link
}
