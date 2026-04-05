package main

import (
	"fmt"
	"testing"
)

type transformFn func(int) int

type anotherFn func(int, []string, map[string][]int) ([]int, string)

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func TestMakeArray(t *testing.T) {
	numbers := []int{1, 23, 34, 5}
	moreNumbers := []int{5, 1, 2}
	double := transformNumbers(&numbers, double)
	triple := transformNumbers(&numbers, triple)

	fmt.Println("numbers :", numbers)
	fmt.Println("double :", double)
	fmt.Println("triple :", triple)

	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
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
