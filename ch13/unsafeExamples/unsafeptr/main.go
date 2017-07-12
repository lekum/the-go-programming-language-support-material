package main

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
}

func main() {
	// equivalent to pb := &x.b
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"
}
