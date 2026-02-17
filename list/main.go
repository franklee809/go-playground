package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [1]string = [1]string{"test123"}

	prices := [5]float64{10.99, 40.9, 12, 12, 123}

	fmt.Println(prices)
	// productNames[0] = "!23"
	fmt.Println(productNames)

	featuredPrices := prices[1:3]

	fmt.Println(featuredPrices)

}
