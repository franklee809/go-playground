package main

import "fmt"
func main() {
	age := 32

	fmt.Println("Age : ", age)
	adultYears := getAdultYears(&age)
	fmt.Println("Adult Years: ", adultYears)
}

func getAdultYears(age *int) int {
	fmt.Println("AGE pointer address : ", age)
	fmt.Println("AGE pointer data: ", *age)
	*age = *age - 18 
	return *age - 18 
} 