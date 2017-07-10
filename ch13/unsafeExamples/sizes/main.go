package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(float64(0))) // 8
	fmt.Println(unsafe.Sizeof(struct {
		bool
		float64
		int16
	}{}))
	fmt.Println(unsafe.Sizeof(struct {
		float64
		int16
		bool
	}{}))
	fmt.Println(unsafe.Sizeof(struct {
		bool
		int16
		float64
	}{}))

}
