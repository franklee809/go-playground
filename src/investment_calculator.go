package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount int = 100
	expectedReturnRate := 5.5

	inflationRate := 1.1
	var years float64 = 10.0
	var futureValue float64 = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %1.f \n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for Inflation): %.2f \n", futureRealValue)

	// fmt.Printf("future Value : %.2f \n Future Value (Adjusted for Inflation): %.2f \n", futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}
