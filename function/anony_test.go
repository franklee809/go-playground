package main

import (
	"fmt"
	"testing"
)

func TestAnonyFunction(t *testing.T) {
	numbers := []int{1, 2, 3}

	double := createTransfomer(2)
	triple := createTransfomer(3)

	transformed := transformNumbers(&numbers, func(number int) int { return number * 2 })

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func createTransfomer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
