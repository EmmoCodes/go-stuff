package practice

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	hobbies := [3]string{"coding", "gaming", "cooking"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])

	hobbiesNew := hobbies[1:]
	fmt.Println(hobbiesNew)

	hobbieSlice := []string{}
	hobbieSlice = hobbies[0:2]
	fmt.Println(hobbieSlice)

	hobbieSlice = hobbies[1:]

	product := []Product{
		Product{"first-product", 0, 599.99},
		Product{"second-product", 1, 299.99},
	}

	newProduct := Product{"third-product", 2, 299.99}

	product = append(product, newProduct)
	fmt.Println(product)

	prices := []float64{10.99, 8.99}
	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
