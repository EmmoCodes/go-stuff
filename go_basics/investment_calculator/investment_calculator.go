package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5

	var investmentAmount, expectedReturnRate, years float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate)

	formattedFV :=

		fmt.Println(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return

}
