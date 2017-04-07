package main

import "fmt"

type Counter struct {
	n int
}

func (c *Counter) N() int {
	return c.n
}

func (c *Counter) Increment() {
	c.n++
}

func (c *Counter) Decrement() {
	c.n--
}

func main() {
	x := Counter{1}
	fmt.Println(x.N())
	x.Increment()
	fmt.Println(x.N())
	x.Decrement()
	fmt.Println(x.N())
}
