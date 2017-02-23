package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	decimal := ""
	sign := ""
	if s[0] == '-' {
		sign = "-"
		s = s[1:]
	}

	if decimalSeparator := strings.LastIndex(s, "."); decimalSeparator > 0 {
		decimal = "." + s[decimalSeparator+1:]
		s = s[:decimalSeparator]

	}

	n := len(s)
	if n <= 3 {
		return sign + s + decimal
	} else {
		return sign + comma(s[:n-3]) + "," + s[n-3:] + decimal
	}
}

func main() {
	fmt.Println(comma("1000"))          // "1,000"
	fmt.Println(comma("10"))            // "10"
	fmt.Println(comma("1000.25"))       // "1,000.25"
	fmt.Println(comma("-1000.25"))      // "-1,000.25"
	fmt.Println(comma("1000000"))       // "1,000,000"
	fmt.Println(comma("-10000000.625")) // "-10,000,000.625"
}
