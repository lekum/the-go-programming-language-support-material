package main

import (
	"fmt"

	"github.com/lekum/the-go-programming-language-support-material/ch6/geometry"
)

func main() {
	p := geometry.Point{2, 3}
	q := geometry.Point{0, 2}
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
	distance := geometry.Point.Distance
	fmt.Println(distance(p, q))
}
