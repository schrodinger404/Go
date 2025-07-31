package main

import "fmt"

// you cannot use this syntax here
// status:="Learning Go"

func main() {

	// variable declare and assignment with data type
	var name string = "Shreekant"
	fmt.Println(name)

	// you can also declare and assign without datatype - Go will infer the datatype internally
	// one thing to notice is once you declare a variable you must use it
	var role = "Engineer"
	fmt.Println(role)

	//short hand - Note : go do not allow using short had syntax outside func, eg. above
	place := "Ingalagi"
	fmt.Println(place)

	// just declare - if you just want to declare you much write full syntax
	var state string
	state = "Karnataka"
	fmt.Println(state)

	age := 27
	fmt.Println(age)

	var isAdult bool = true;
	fmt.Println("is adult", isAdult)

	var price float32 = 50.5;
	fmt.Println(price/2)

}