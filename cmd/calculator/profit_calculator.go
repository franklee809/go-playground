package main

import "fmt"

func main() {
	// Revenue
	revenue := handleUserInput("Enter revenue: ")
	expenses := handleUserInput("Enter Expense: ")
	tax_rate := handleUserInput("Enter Tax rate: ")

	ebt_value, earning_after_tax, ratio := calculateFinancials(revenue, expenses, tax_rate)
	fmt.Printf("EBT : %.1f \n", ebt_value)
	fmt.Printf("Profit : %.1f \n", earning_after_tax)
	fmt.Printf("Ratio : %.12f \n", ratio)

}
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt_value := revenue - expenses
	tax_expenses := float64(ebt_value*taxRate) / 100
	earning_after_tax := float64(ebt_value) - tax_expenses

	ratio := float64(ebt_value) / float64(earning_after_tax)

	return ebt_value, earning_after_tax, ratio
}

func handleUserInput(infoText string) float64 {
	input := 0.0
	// Revenue
	print(infoText)

	fmt.Scan(&input)

	return input
}
