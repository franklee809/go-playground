package main

import (
	"fmt"
	"time"
)

type user struct {
	lastName  string
	firstName string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	fmt.Println("User Details:")
	fmt.Printf("First Name: %s\n", u.firstName)
	fmt.Printf("Last Name: %s\n", u.lastName)
	fmt.Printf("Birth Date: %s\n", u.birthDate)
	fmt.Printf("Account Created At: %s\n", u.createdAt.Format(time.RFC1123))
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	firstName := getUserData("Please enter your first name : ")
	lastName := getUserData("Please enter your last name : ")
	birthDate := getUserData("Please enter your birth date (DD/MM/YYYY) : ")

	var appUser user
	appUser = user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	var value string
	fmt.Scan(&value)
	return value
}
