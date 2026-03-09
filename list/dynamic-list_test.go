package main

import (
	"fmt"
	"testing"
)

func TestDynamicList(t *testing.T) {
	prices := []float64{10.99, 92.21}
	fmt.Println(prices[0:1])
	prices[1] = 91.21
	discountPrices := []float64{102.213, 12.23, 232.23}
	prices = append(prices, discountPrices...)
	// fmt.Println(prices[:1])
	fmt.Println(prices)

}
