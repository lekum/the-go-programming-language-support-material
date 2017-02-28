package main

import "fmt"

const (
	_ = 1 << (10 * iota) * 8
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	powers := []int{KB, MB, GB}
	for i, pow := range powers {
		fmt.Println(i+1, "->", pow)
	}
}
