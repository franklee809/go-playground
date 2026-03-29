package main

import (
	"fmt"
	"testing"
)

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
	fmt.Println(userNames)

	courseRatings := make(map[string]float64)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 5.8
	fmt.Println(courseRatings)
}
