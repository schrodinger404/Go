package main

import (
	"fmt"
	"slices"
)

func main() {

// slices are dynamic arrays
// when you declare a slice with zero elements, it'll be nill, and has length zero
	var s []int
	fmt.Println("Unintialized slice:",s, len(s), s==nil)

// use builtin make func to declare a non nill slice, whos lenght and capacity will be 3 in below example 
	names := make([]string,3)
	fmt.Println("Slice with builtin make func:", names, len(names), cap(names), names == nil)


// declare slice with initial length 0 and capacity 5 and 
	names = make([]string,0,5)
	fmt.Println("Slice with builtin make func:", names, len(names), cap(names), names == nil)

	names = append(names, "Ben")
	names = append(names, "Jon")
	names = append(names, "Chris")
	names = append(names, "shrrknt")
	names = append(names, "Yash")
	names = append(names, "Jack")

	fmt.Println("Slice with builtin make func:", names, len(names), cap(names), names == nil)
	// output => Slice with builtin make func: [Ben Jon Chris shrrknt Yash Jack] 6 10 false
	// observe how slice has increased length and capacity

	c := make([]string, len(names))
    copy(c, names)
    fmt.Println("copy:", c, len(c), cap(c))
	c = append(c, "Jacki")
	fmt.Println("append Jacki:", c, len(c), cap(c))
	// output => copy: [Ben Jon Chris shrrknt Yash Jack Jacki] 7 12

// making an slice with c
	m := c[2:5]
	fmt.Println("slice m made from c[2:5]",m, len(m), cap(m))

	m = c[:3]
	fmt.Println("slice m made from c[:3]",m, len(m), cap(m))

	m = c[3:]
	fmt.Println("slice m made from c[3:]",m, len(m), cap(m))

	fmt.Println(slices.Equal(names,c)) //false


// you can declare and assign values to slice 
	t := []string{"a", "b", "c"}
    fmt.Println("slice t:", t)

	p :=[]string{"A", "b", "c"}
	fmt.Println("slice p: ", p)

	fmt.Println(slices.Equal(t,p))

}