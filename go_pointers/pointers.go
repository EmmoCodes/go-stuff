package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", *agePointer)

	getAdultYears(&age)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
