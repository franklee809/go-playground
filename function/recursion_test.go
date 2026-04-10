package main

import (
	"fmt"
	"testing"
)

func TestRecursion(t *testing.T) {
	fmt.Println(factorial(5))
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	result := number * factorial(number-1)
	return result
	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }
	// return result
}
