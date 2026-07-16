package main

import (
	"fmt"
	"testing"

	"example.com/practice_function/filemanager"
	"example.com/practice_function/prices"
)

func TestTaxRate(t *testing.T) {
	var taxRates []float64 = []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanagergo.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}

	}
}

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}
