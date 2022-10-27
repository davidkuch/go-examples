package variadic

import "fmt"

func v1() {

	sum := variadic_sum(1, 2, 3, 4, 5, 6, 7, 8)

	fmt.Println("sum is: ", sum)
}

func variadic_sum(nums ...int) int {
	var ret int

	for num := range nums {
		ret += num
	}

	return ret
}
