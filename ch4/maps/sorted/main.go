package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"Martina":   21,
		"Alejandro": 13,
		"Estefanía": 4,
		"Marcos":    52,
		"Irene":     22,
	}
	var names []string
	for name, _ := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name, "->", ages[name])
	}
}
