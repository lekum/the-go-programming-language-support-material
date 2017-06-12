package main

import (
	"fmt"
	"github.com/lekum/the-go-programming-language-support-material/ch9/memo1"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() []string {
	return []string{
		"https://lekum.org",
		"https://gnupg.org",
		"https://lekum.org",
	}
}

func main() {
	m := memo.New(httpGetBody)
	for _, url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}
