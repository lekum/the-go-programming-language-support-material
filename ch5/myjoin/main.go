package main

import (
	"fmt"
	"os"
	"strings"
)

func myJoin(sep string, l ...string) string {
	return strings.Join(l, sep)
}

func main() {
	var sep string
	if len(os.Args) < 2 {
		sep = " "
	} else {
		sep = os.Args[1]
	}
	fmt.Println(myJoin(sep, "Hi", "my", "name", "is", "Gordon"))
}
