package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount int = 100
	expectedReturnRate := 5.5

	var years int = 10

	var futureValue float64 = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	test := 01232.0
	fmt.Printf("future Value : %.2f \n Future Value (Adjusted for Inflation): %.2f \n", futureValue, test)

}
