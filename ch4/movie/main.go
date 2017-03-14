package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title:  "Casablanca",
		Year:   1942,
		Color:  false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title:  "Cool Hand Luke",
		Year:   1967,
		Color:  true,
		Actors: []string{"Steve McQueen", "Jackeline Bisset"},
	},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		fmt.Printf("JSON marshaling failed: %s\n", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	err = json.Unmarshal(data, &titles)
	if err != nil {
		fmt.Printf("JSON unmarshaling failed: %s\n", err)
	}
	for _, title := range titles {
		fmt.Println(title.Title)
	}
}
