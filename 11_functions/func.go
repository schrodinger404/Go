package main

import "fmt"

// single reuturn
func add(a, b int) int {
	return a + b

}

// multiple return
func multipleReturn()  (string, int, bool){
	return "shrrknt", 27, true
	
}

// functions can also take input as function
func funcInput(num func(s int) int ) {
	num(10)
}

func main() {
	res := add(45, 5)
	fmt.Println(res)

	fmt.Println(multipleReturn())

	x := func (i int) int {
		return 20;
		
	}

	funcInput(x)
	fmt.Println()


}

// need to comeback understand in much details