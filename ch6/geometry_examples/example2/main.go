package main

import (
	"fmt"

	"github.com/lekum/the-go-programming-language-support-material/ch6/geometry"
)

func main() {
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Printf("perim: %v\n", perim)
	fmt.Println("Distance:", perim.Distance())

}
