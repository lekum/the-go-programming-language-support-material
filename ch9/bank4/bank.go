// Package bank provides a concurrency-safe bank with one account
package bank

import "sync"

var (
	mu      sync.RWMutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.RLock() // readers lock
	defer mu.RUnlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock() // exclusive lock
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}

// This function requeres that the lock be held.
func deposit(amount int) { balance += amount }
