package main

import (
	"fmt"

	cmdmanagergo "example.com/practice_function/cmdmanager.go"
	"example.com/practice_function/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		cmdm := cmdmanagergo.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)

		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			fmt.Println(err)
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
