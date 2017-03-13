package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w1, w2 Wheel
	w1 = Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v\n", w1)
	w2 = Wheel{
		Circle: Circle{
			Point:  Point{X: 7},
			Radius: 5,
		},
		Spokes: 15,
	}
	fmt.Printf("%#v\n", w2)
}
