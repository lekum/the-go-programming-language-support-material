package main

import (
	"fmt"
	"unsafe"
)

func Float64bits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }

func main() {
	fmt.Printf("%#016x\n", Float64bits(1.0)) // "0x3ff0000000000000"
}
