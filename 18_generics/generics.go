package main

import "fmt"


// this is restricted to only receive slice of integers
func print(vals []int) {

	for _, v := range vals {

		fmt.Println(v)

	}
}

// using generics, this func will receive any type
func printAny[T any](vals []T)  {

	for _, v := range vals {

		fmt.Println(v)
		
	}
}

// this func can receive only string, int or bool
func printRestricted [T string|int|bool] (vals []T)  {
	for _, v := range vals {
		fmt.Println(v)
	}
}

// we can make a generic struct also
type something[T any] struct {
	elements []T
	
}

func main() {
	nums:=[]int{10,40,35,67}
	// function calls
	print(nums)
	printAny(nums)

	langs:=[]string{"go","java"}
	printAny(langs)

	boo:= []bool{true, false,false, true}
	printRestricted(boo)

	// instanstiating something struct
	mySomething:= something[string]{
		elements: []string{"anything","something", "everything"},
	}

	fmt.Println(mySomething)
}