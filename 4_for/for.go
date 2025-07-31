package main

import "fmt"

func main() {

	// dont hav while loop in Go, so you can youse something like this to achieve while loop
	i := 0
	for i < 5 {

		fmt.Println(i)
		i ++

	}

	// classic for loop
	for i := 0; i <5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
		
	}

}