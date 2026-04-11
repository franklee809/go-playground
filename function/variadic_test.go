package main

import (
	"fmt"
	"testing"
)

func TestVariadic(t *testing.T) {
	// numbers := []int{1, 10, 15}
	sum := sumup(1, 2, 3)

	fmt.Println(sum)

}

func sumup(numbers ...int) int {
	fmt.Println(numbers)
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
