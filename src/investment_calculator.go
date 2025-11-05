package main

import (
	"math"
)

func main() {
	var investmentAmount int = 100
	expectedReturnRate := 5.5

	var years int = 10

	var futureValue float64 = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	println("futureValue : ", futureValue)

}
