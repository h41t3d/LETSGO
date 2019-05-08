package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	gg := max(a, b)
	fmt.Printf("aha %d", gg)
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
