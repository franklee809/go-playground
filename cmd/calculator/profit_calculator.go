package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// Revenue
	revenue, err := handleUserInput("Enter revenue: ")

	if err != nil {
		panic(err)
	}
	expenses, err := handleUserInput("Enter Expense: ")
	if err != nil {
		panic(err)
	}
	tax_rate, err := handleUserInput("Enter Tax rate: ")

	if err != nil {
		panic(err)
	}

	ebt_value, earning_after_tax, ratio := calculateFinancials(revenue, expenses, tax_rate)
	fmt.Printf("EBT : %.1f \n", ebt_value)
	fmt.Printf("Profit : %.1f \n", earning_after_tax)
	fmt.Printf("Ratio : %.12f \n", ratio)
	writeBalanceToFile("ebt_value", ebt_value)
	writeBalanceToFile("earning_after_tax", earning_after_tax)
	writeBalanceToFile("ratio", ratio)
}
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt_value := revenue - expenses
	tax_expenses := float64(ebt_value*taxRate) / 100
	earning_after_tax := float64(ebt_value) - tax_expenses

	ratio := float64(ebt_value) / float64(earning_after_tax)

	return ebt_value, earning_after_tax, ratio
}

func handleUserInput(infoText string) (float64, error) {
	input := 0.0
	// Revenue
	print(infoText)

	fmt.Scan(&input)
	if input <= 0 {
		return 0, errors.New("value must be positive")
	}
	return input, nil
}

func writeBalanceToFile(name string, amount float64) {
	var amountString string = fmt.Sprint(amount)
	os.WriteFile(name, []byte(amountString), os.FileMode(0644))
}
