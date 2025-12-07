package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func readBalanceFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000.0, errors.New("failed to find new balance file")
	}
	fmt.Println("data :", string(data))
	value, err := strconv.ParseFloat(string(data), 32)

	if err != nil {
		return 1000.0, errors.New("failed to parse value")
	}
	return value, nil
}

func writeBalanceToFile(balance float64, filename string) {
	var balanceText string = fmt.Sprint(balance)
	fmt.Println(balanceText)
	os.WriteFile(filename, []byte(balanceText), os.FileMode(0644))
}

func main() {
	balance, err := readBalanceFromFile("test.txt")

	if err != nil {
		fmt.Println("Error")
		fmt.Println("The error value : ", err)
		fmt.Println("------------------------------------------------------------------------------------------")
		panic("Can't continue, sorry")
	}
	var accountBalance = balance
	writeBalanceToFile(accountBalance, "test.txt")

	fmt.Println("Welcome to the Go Bank!")

	for {
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
			os.Exit(0)
			// return
		}

	}

}
