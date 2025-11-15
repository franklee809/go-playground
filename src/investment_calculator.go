package main

import (
	"fmt"
	"math"
)

const inflationRate = 1.1

func main() {
	investmentAmount := 100.0
	expectedReturnRate := 5.5

	var years float64 = 10.0

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formatText(futureValue, futureRealValue)
}

func formatText(futureValue float64, futureRealValue float64) {
	// fmt.Printf("future Value : %.2f \n Future Value (Adjusted for Inflation): %.2f \n", futureValue, futureRealValue)
	fmt.Sprintf("Future Value: %v \n", futureValue)
	fmt.Sprintf("Future Value (Adjusted for Inflation):%.2f \n", futureRealValue)
	print()
}

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	return
}
