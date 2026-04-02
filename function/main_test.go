package main

import (
	"fmt"
	"testing"
)

type transformFn func(int) int

func TestMakeArray(t *testing.T) {
	numbers := []int{1, 23, 34, 5}
	double := doubleNumbers(&numbers, double)
	triple := doubleNumbers(&numbers, triple)

	fmt.Println("numbers :", numbers)
	fmt.Println("double :", double)
	fmt.Println("triple :", triple)
}

func doubleNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		fmt.Println(value)
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
