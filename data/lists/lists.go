package data

import "fmt"

func lists() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)

	updatedPrices := append(prices, 5.99, 12.99, 29.99, 199.00)
	fmt.Println(updatedPrices)

	discountPrices := []float64{101.99, 90.99, 20.95}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames = [4]string{"A Book"}
// 	prices := [4]float64{1.99, 2.99, 5.43, 20.0}
// 	fmt.Println(prices)
//
// 	productNames[2] = "A Carpet"
//
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])
//
// 	featuredPrices := prices[1:]
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
//
// }
