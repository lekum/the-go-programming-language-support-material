package main

import (
	"fmt"

	"github.com/lekum/the-go-programming-language-support-material/ch6/geometry"
)

func main() {
	p := geometry.Point{2, 3}
	p.ScaleBy(2.5)
	fmt.Println("p:", p)
}
