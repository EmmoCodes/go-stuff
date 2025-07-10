package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

func add[T float64 | int | string](a, b T) T {
	return a + b
}
