package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple switch
	i := 2

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")

	default:
		fmt.Print("Unknown")

	}

	// multiple condition switch
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("Weekend")
	

		default:
			fmt.Println("Week day")
			
	}

	// type switch
	WhatIsIt:= func(i interface{})  {
		switch i.(type){
		case int:
			fmt.Println("Integer")
		case bool:
			fmt.Println("Boolean")
		case string:
			fmt.Println("String")
		default:
			fmt.Println("Unknown")
		}
		
	}
	WhatIsIt(10)

}