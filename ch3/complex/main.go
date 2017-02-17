package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Println(1i * 1i)
	fmt.Println(cmplx.Sqrt(-1))
	a := 2 + 3i
	fmt.Println(a)
	fmt.Println(cmplx.Pow(a, 1.0/2.0))
	fmt.Println(cmplx.Sqrt(a))
}
