package main

import "fmt"

// This function seq returns another function, which we define anonymously in the body of seq. 
// The returned function closes over the variable i to form a closure.
func seq() func() int {

	i := 0

	return func() int {
		i++
		return i
	}

}

func main() {

	//We call seq, assigning the result (a function) to val. 
	// This function value captures its own i value, which will be updated each time we call val.
	val := seq()

	// See the effect of the closure by calling nextInt a few times.
	fmt.Println(val()) // o/p: 1
	fmt.Println(val()) // o/p: 2


	// To confirm that the state is unique to that particular function, create and test a new one.
	res := seq()

	fmt.Println(res()) // o/p: 1

}