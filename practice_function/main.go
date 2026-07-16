package main

import (
	cmdmanagergo "example.com/practice_function/cmdmanager.go"
	"example.com/practice_function/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		cmdm := cmdmanagergo.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}
}
