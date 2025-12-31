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


func main(){
	firstName := getUserData("Please enter your first name : ")
	lastName := getUserData("Please enter your last name : ")
	birthDate := getUserData("Please enter your birth date (DD/MM/YYYY) : ")

	var appUser user
	appUser = user {
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
	fmt.Println(appUser)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	var value string;
	fmt.Scan(&value)
	return value
}