package main

import "fmt"

func main() {

// array with 5 int values, if not initialized then by default will be zero
	var i [5]int
	i[0] = 10

	fmt.Println(i)
	fmt.Println(len(i))
	fmt.Println(cap(i))

	j:=[4]string{"shreekant","teja"}
	fmt.Println(j)
	fmt.Println(len(j))

	// You can also have the compiler count the number of elements for you with ...
    b := [...]int{1, 2, 3, 4, 5}
    fmt.Println(b)

	//If you specify the index with :, the elements in between will be zeroed.
	c := [...]int{1, 3: 4, 5}
    fmt.Println(c)

}