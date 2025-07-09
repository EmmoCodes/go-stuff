package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()

	for i, val := range userNames {
		fmt.Println("Index: ", i)
		fmt.Println("Val: ", val)
	}

	for key, val := range courseRatings {
		fmt.Println("Key: ", key)
		fmt.Println("Val : ", val)
	}
}
