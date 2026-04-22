package main

import (
	"testing"

	"example.com/practice_function/prices"
)

func TestTaxRate(t *testing.T) {
	var taxRates []float64 = []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()

	}
}
