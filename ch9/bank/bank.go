package main

import (
	"fmt"
	"sync"
)

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

func main() {
	for i := 0; i < 100; i++ {
		// Alice:
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			Deposit(200)
			fmt.Println("=", Balance())
			wg.Done()
		}()
		// Bob:
		go func() {
			Deposit(100)
			wg.Done()
		}()
		wg.Wait()
	}
}
