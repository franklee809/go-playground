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
	fmt.Println(myHobbies[:2])

	//q4
	fmt.Println(cap(myHobbies))
	fmt.Println(myHobbies[1:3])

	//q5
	courseGoal := []string{"course goal 1"}
	fmt.Println(courseGoal)
	//q6
	courseGoal = append(courseGoal, "course goal 2")
	courseGoal = append(courseGoal, "course goal 3")
	//q7
	products := []ProductTwo{
		ProductTwo{"first-product", "A first product", 12.21},
	}
	fmt.Println(products)
	newProducts := ProductTwo{"third-product", "third", 123123}
	products = append(products, newProducts)
	fmt.Println(products)
}

// q7

type ProductTwo struct {
	title string
	id    string
	price float64
}
