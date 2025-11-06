package main

import "fmt"

func main() {
	revenue := 0
	expenses := 0
	tax_rate := 0.0
	println("hello")
	// Revenue
	print("Enter revenue: ")

	fmt.Scan(&revenue)
	print("Enter Expense: ")
	// Expense
	fmt.Scan(&expenses)

	print("Enter Tax rate: ")

	// Tax rate
	fmt.Scan(&tax_rate)

	ebt_value := revenue - expenses
	tax_expenses := float64(ebt_value * int(tax_rate) / 100)
	earning_after_tax := float64(ebt_value) - tax_expenses

	ratio := float64(ebt_value) / float64(earning_after_tax)

	print("EBT : ", ebt_value, "\n")
	print("Profit : ", int(earning_after_tax), "\n")
	print("Ratio : ")
	fmt.Printf("%.12f\n", ratio)

}
