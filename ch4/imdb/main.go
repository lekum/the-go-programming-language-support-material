package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Movie struct {
	Title  string
	Poster string
}

func main() {
	title := os.Args[1]
	url := "http://www.omdbapi.com/"
	res, err := http.Get(url + "?t=" + title)
	if err != nil {
		fmt.Println("Error contacting IMDB API:", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	var m Movie
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		os.Exit(1)
	}
	fmt.Println(m.Title)
	fmt.Println(m.Poster)
}
