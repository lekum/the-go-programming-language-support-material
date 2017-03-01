package main

import "fmt"

func main() {
	type Currency int

	const (
		USD Currency = iota
		EUR
		ZMB
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", RMB: "w00t!", GBP: "pound"}
	fmt.Println(symbol)
}
