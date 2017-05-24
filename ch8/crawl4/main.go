package main

import (
	"flag"
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"strings"
)

// tokens is a counting semaphore used to enforce
// a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)

// worklist of urls list (may contain duplicates)
type worklist struct {
	currentDepth int
	urls         []string
}

// unseen url
type unseenLink struct {
	currentDepth int
	url          string
}

func crawl(url string) []string {
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print("Error extracting links: ", err)
	}
	return list
}

func main() {
	var maxDepth = flag.Int("depth", 100, "max depth of crawling")
	flag.Parse()
	wl := make(chan worklist)
	ul := make(chan unseenLink)

	// Add command-line arguments to worklist.
	go func() {
		wl <- worklist{0, flag.Args()} // currentDepth=0
	}()

	// Create 20 crawler goroutines to fetch each unseen link
	for i := 0; i < 20; i++ {
		go func() {
			for link := range ul {
				foundLinks := crawl(link.url)
				go func() {
					wl <- worklist{link.currentDepth + 1, foundLinks}
				}()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range wl {
		if list.currentDepth >= *maxDepth {
			// maxDepth reached, ignoring list
			continue
		}
		for _, link := range list.urls {
			if !seen[link] {
				seen[link] = true
				fmt.Printf("%s%s\n", strings.Repeat("\t", list.currentDepth), link)
				ul <- unseenLink{list.currentDepth, link}
			}
		}
	}
}
