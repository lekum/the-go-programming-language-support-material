package main

import (
	"bytes"
	"fmt"
)

func printints(values []int) string {
	var b bytes.Buffer
	b.WriteByte('[')
	for i, val := range values {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%d", val)
	}
	b.WriteByte(']')
	return b.String()

}

func main() {
	fmt.Println(printints([]int{1, 2, 3}))
}
