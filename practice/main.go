package main

import "fmt"

func main() {

	hobbies := []string{"coding", "gaming", "nature"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	updatedHobbies := hobbies[0:2]
	hobbies = hobbies[0:2]
	fmt.Println(updatedHobbies)
	fmt.Println(hobbies)
	updatedHobbies = hobbies[1:3]
	fmt.Println(updatedHobbies)

	dynArr := []string{"master Golang", "write clean go code"}
	dynArr[1] = "write cleaner go code!"
	dynArr = append(dynArr, "become software dev #noJS")

	fmt.Println(dynArr)

	dynProducts := []products{}
	dynProducts = append(dynProducts, products{title: "Apple", id: "1", price: 5.99})
	dynProducts = append(dynProducts, products{title: "Orange", id: "2", price: 2.99})
	fmt.Println(dynProducts)
}

type products struct {
	title string
	id    string
	price float64
}

// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
