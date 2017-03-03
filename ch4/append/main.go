package main

import "fmt"

func appendInt(x []int, v int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = v
	return z
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(appendInt(a, 4))
}
