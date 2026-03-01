package main

import (
	"fmt"
	"testing"
)

func TestPractice(t *testing.T) {
	myHobbies := []string{"swimming", "gymnastic", "cafe hopping"}
	fmt.Println(myHobbies)
	// q2
	fmt.Println(myHobbies[0])
	fmt.Println(myHobbies[1:])

	//q3
	fmt.Println(myHobbies[0:2])

	//q4
	fmt.Println(myHobbies[2:])
	fmt.Println(myHobbies[1:])

	//q5
	courseGoal := []string{"course goal 1"}
	fmt.Println(courseGoal)
	//q6
	courseGoal = append(courseGoal, "course goal 2")
	courseGoal = append(courseGoal, "course goal 3")

}

// q7

type ProductTwo struct {
	title string
	id    string
	price float64
}
