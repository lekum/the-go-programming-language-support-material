package main

import (
	"fmt"

	"github.com/lekum/the-go-programming-language-support-material/ch6/geometry"
)

func main() {
	p := geometry.Point{0, 0}
	q := geometry.Point{1, 1}
	fmt.Printf("p: %v\n", p)
	fmt.Printf("q: %v\n", q)
	fmt.Println("Distance:", geometry.Distance(p, q))
	fmt.Println("Distance:", p.Distance(q))
}
