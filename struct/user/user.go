package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	lastName  string
	firstName string
	birthDate string
	createdAt time.Time
}

func (u User) OutputUserDetails() {
	fmt.Println("User Details:")
	fmt.Printf("First Name: %s\n", u.firstName)
	fmt.Printf("Last Name: %s\n", u.lastName)
	fmt.Printf("Birth Date: %s\n", u.birthDate)
	fmt.Printf("Account Created At: %s\n", u.createdAt.Format(time.RFC1123))
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
