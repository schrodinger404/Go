package main

import "fmt"

// pass by value
func passByVal(i int)  {
	i = 2
	fmt.Println("value in passByVal:",i)
}

func passByRef(i *int)  {
	*i = 3
	fmt.Println("value in passByRef:", *i)
	
}

func main() {

	num := 0

	fmt.Println("original num in main:",num)
	fmt.Println("original ref in main:",&num)

	passByVal(num) // this is in the passByVal function
	fmt.Println("value in main after passing it as value:",num) // this is the value in main func

	passByRef(&num)
	fmt.Println("value in main after passing as ref:",num)
	
}