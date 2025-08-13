package main

import "fmt"

func sum(num ...int) int{

	total := 0

	for _, v := range num {
		total += v

	}
	return total
}

func main() {

	// you can pass number of elements if you have variadic functions
	value := sum(100, 45)

	fmt.Println(value)


	// you can also pass slices as shown here
	nums := []int {12, 43, 68,98}
	res := sum(nums...)

	fmt.Println(res)

}