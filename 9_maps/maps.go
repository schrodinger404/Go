package main

import (
	"fmt"
 	"maps"
)

func main() {
	// declaring a map
	m := make(map[int]string)

	// assign values
	m[1] = "shrrknt"
	m[2] = "teja"
	m[3] = "someone"

	fmt.Println(m)

	// retrive values
	fmt.Println("m[2] value:",m[2])

	// if key does not exist then returns correcsponding default values
	fmt.Println("m[4] value:",m[4]) // this will return empty space,

	fmt.Println(len(m))

	// delete 3 key value pair
	delete(m, 3)
	fmt.Println(m)

	// declare and assign 
	m1:= map[string]int{"age":27, "salary":464000}
	fmt.Println(m1)

	// here v will return value
	v, ok := m1["salary"]
	fmt.Println(v)

	if ok {
		fmt.Println("all ok")
	}else {
		fmt.Println("not ok")
	}



	// builtin equals function
	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}