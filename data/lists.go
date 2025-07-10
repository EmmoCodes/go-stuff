package main

import "fmt"

func main() {
	var productNames = [4]string{"A Book"}
	prices := [4]float64{1.99, 2.99, 5.43, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)
	fmt.Println(prices[2])

	featuredPrices := prices[1:]
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))

}
