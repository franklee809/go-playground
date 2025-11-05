package main

import "fmt"

func main() {
	revenue := 0
	expenses := 0
	tax_rate := 0.0
	println("hello")
	// Revenue
	println("Enter revenue: ")

	fmt.Scan(&revenue)
	println("Enter Expense: ")
	// Expense
	fmt.Scan(&expenses)

	println("Enter Tax rate: ")

	// Tax rate
	fmt.Scan(&tax_rate)

	ebt_value := revenue - expenses
	tax_expenses := float64(ebt_value * int(tax_rate) / 100)
	earning_after_tax := float64(ebt_value) - tax_expenses

	ratio := float64(ebt_value) / float64(earning_after_tax)
	// The basic formula for calculating EBT is: \(\text{EBT}=\text{Total\ Revenue}-\text{Cost\ of\ Goods\ Sold\ (COGS)}-\text{Operating\ Expenses}-\text{Interest\ Expense}\)Alternatively, you can start with a company's pre-calculated EBIT (Earnings Before Interest and Taxes) and simply subtract the interest expense:
	// Output EBT, profit and the ratio

	print("EBT : ", ebt_value, "\n")
	print("Profit : ", int(earning_after_tax), "\n")
	print("Ratio : ")
	fmt.Printf("%.12f\n", ratio)

}
