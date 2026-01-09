package main

import (
	"fmt"

	"example.com/struct/user"
)

func main() {
	firstName := getUserData("Please enter your first name : ")
	lastName := getUserData("Please enter your last name : ")
	birthDate := getUserData("Please enter your birth date (DD/MM/YYYY) : ")

	var appUser *user.User
	appUser, err := user.NewUser(firstName, lastName, birthDate)

	if err != nil {
		println(err.Error())
		return
	}
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	var value string
	fmt.Scanln(&value)
	return value
}
