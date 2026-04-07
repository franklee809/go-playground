package main

import (
	"fmt"
	"testing"
)

func TestAnonyFunction(t *testing.T) {
	numbers := []int{1, 2, 3}

	fmt.Println("numbers :", numbers)

	transformed := transformNumbers(&numbers, func(number int) int { return number * 2 })

	fmt.Println(transformed)
}
