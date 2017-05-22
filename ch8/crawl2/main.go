package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

// tokens is a counting semaphore used to enforce
// a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// n is the numbers of pending sends to worklist
	var n int
	worklist := make(chan []string)

	// Start with the command-line arguments
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	// Crawl the web concurrently
	for ; n > 0; n-- {
		seen := make(map[string]bool)
		for list := range worklist {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
		}
	}
}
