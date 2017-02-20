package main

import "fmt"

func main() {
	greeting := "Hello, \xe4\xb8\x96\xe7\x95\x8c"
	fmt.Println(greeting)
	for i, r := range greeting {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
