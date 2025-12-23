package main

import (
	"fmt"
	"os"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	balance, err := fileops.ReadBalanceFromFile("test.txt")

	if err != nil {
		fmt.Println("Error")
		fmt.Println("The error value : ", err)
		panic("Can't continue, sorry")
	}

	var accountBalance = balance
	fileops.WriteBalanceToFile(accountBalance, "test.txt")

	fmt.Println("Welcome to the Go Bank!")
	fmt.Println("Phone: ", randomdata.PhoneNumber())
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
