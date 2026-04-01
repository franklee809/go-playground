package main

import (
	"fmt"
	"testing"
)

// type aliases
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func TestMakeArray(t *testing.T) {
	// for memory managements
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")
	// fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 5.8
	courseRatings["angular"] = 5.12
	// fmt.Println(courseRatings)
	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}

	for key, value := range courseRatings {
		fmt.Println("key : ", key)
		fmt.Println("value : ", value)
	}
}
