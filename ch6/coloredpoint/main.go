package main

import (
	"fmt"
	"image/color"

	"github.com/lekum/the-go-programming-language-support-material/ch6/geometry"
)

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{geometry.Point{1, 1}, red}
	var q = ColoredPoint{geometry.Point{5, 4}, blue}
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p, q)
	fmt.Println(p.Distance(q.Point))
}
