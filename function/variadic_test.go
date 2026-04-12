package main

import (
	"fmt"
	"testing"
)

func TestVariadic(t *testing.T) {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 2, 3)

	anotherSum := sumup(1, numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)

}

func sumup(startingValue int, numbers ...int) int {
	fmt.Println(numbers)
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
