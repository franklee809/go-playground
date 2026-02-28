package main

import (
	"fmt"
	"testing"
)

func TestDynamicList(t *testing.T) {
	prices := []float64{10.99, 92.21}
	fmt.Println(prices[0:1])
	prices[1] = 91.21
	prices = append(prices, 191.21)
	// fmt.Println(prices[:1])
	// fmt.Println(prices)

}
