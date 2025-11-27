package main

import (
	"fmt"
	"os"
	"strconv"
)

func readBalanceFromFile() float64 {
	data, _ := os.ReadFile("test.txt")
	// if err == nil {
	fmt.Println("data :", string(data))
	balance, _ := strconv.ParseFloat(string(data), 32)
	return balance
	// }
}

func writeBalanceToFile(balance float64) {
	var balanceText string = fmt.Sprint(balance)
	fmt.Println(balanceText)
	os.WriteFile("test.txt", []byte(balanceText), os.FileMode(0644))
}

func main() {
	balance := readBalanceFromFile()
	var accountBalance = balance
	writeBalanceToFile(accountBalance)
	return

	fmt.Println("Welcome to the Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	for range 2 {

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1
		if choice == 1 {
			fmt.Println("Your balance is : ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount : ", accountBalance)

		} else if choice == 3 {
			fmt.Print("Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			accountBalance = accountBalance - withdrawalAmount
			fmt.Println("Balance updated! New amount : ", accountBalance)
		} else {
			fmt.Println("Bye bye")
		}

		fmt.Println("Your Choice: ", choice)
	}

}
