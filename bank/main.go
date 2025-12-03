package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		return 1000.0, errors.New("failed to find new balance file")
	}
	fmt.Println("data :", string(data))
	balance, err := strconv.ParseFloat(string(data), 32)

	if err != nil {
		return 1000.0, errors.New("failed to parse balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	var balanceText string = fmt.Sprint(balance)
	fmt.Println(balanceText)
	os.WriteFile("test.txt", []byte(balanceText), os.FileMode(0644))
}

func main() {
	balance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println("The error value : ", err)
		fmt.Println("------------------------------------------------------------------------------------------")
		panic("Can't continue, sorry")
	}
	var accountBalance = balance
	writeBalanceToFile(accountBalance)

	fmt.Println("Welcome to the Go Bank!")

	for range 2 {
		presentOptions()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1
		switch choice {
		case 1:
			fmt.Println("Your balance is : ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount : ", accountBalance)

		case 3:
			fmt.Print("Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			accountBalance = accountBalance - withdrawalAmount
			fmt.Println("Balance updated! New amount : ", accountBalance)
		default:
			fmt.Println("Bye bye")
			fmt.Println("Thanks for choosing our bank")
		}

	}

}
