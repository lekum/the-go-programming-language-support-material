package main

import "fmt"

func sum(args ...float64) float64 {
	total := 0.0
	for _, x := range args {
		total += x
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	a := []float64{1, 2, 3}
	fmt.Println(sum(a...))
}
