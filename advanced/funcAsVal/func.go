package funcasval

import (
	"fmt"
)

type transformFn func(int) int

func funcasval() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	transformerFn1 := getTransformer(&numbers)
	transformerFn2 := getTransformer(&moreNumbers)

	transformedNum := transformNumbers(&numbers, transformerFn1)
	fmt.Println(transformedNum)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNum := []int{}
	for _, val := range *numbers {
		dNum = append(dNum, transform(val))
	}
	return dNum
}

func getTransformer(numbers *[]int) transformFn {

	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}

}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}
