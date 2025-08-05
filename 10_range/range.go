package main

import "fmt"

func main() {
	// iterate over slice
	nums := []int{10, 20, 30, 100}

	// range will always return two values, index and the value incase of slices
	for i, num := range nums {
		fmt.Println(i,num)
	}

	// if i dont have two return here ie instead of i, num if i have only num, this will always return index
	for num := range nums {
		fmt.Println(num)
	}


	sum :=0
	for _, s:= range nums {
		sum += s
	}
	fmt.Println(sum)

	// iterate over map
	m1:= map[string]int{"age":27, "salary":464000}

	// here is is returning key and value, 
	for k,v := range m1{
		fmt.Println(k, v)
	}

	// here it is returning only key 
	for t := range m1{
		fmt.Println(t)
	}

	// iterate over string
	s:= "shreekant"
	//this will print index & unicode
	for i,c:= range s{
		fmt.Println(i,c)
	}


}