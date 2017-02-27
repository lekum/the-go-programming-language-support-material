package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		var integ int
		var err error
		if integ, err = strconv.Atoi(arg); err != nil {
			fmt.Printf("Error processing %s: %s\n", arg, err)
			continue
		}
		fmt.Println("Converted to int:", integ)
		fmt.Printf("Decimal: %d\n", integ)
		fmt.Printf("Binary: %b\n", integ)
	}
}
