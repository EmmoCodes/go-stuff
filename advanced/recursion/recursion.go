package recursion

import "fmt"

func recursion() {

	fact := factorial(6)
	fmt.Println(fact)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)

	//normal for loop
	// result := 1
	//
	// for i := 1; i <= num; i++ {
	// 	result = result * i
	// }
	//
	// return result
}

// factorial of 5 => 5*4*3*2*1 => 120
