package main

import "fmt"

func main() {

	// age := 18

	// if age <= 18 {
	// 	fmt.Println("Adult")
	// }else if (age <=12){
	// 	fmt.Println("Child")
	// }

	if age:=13; age >18{
		fmt.Println("Adult")
	}else if age >=12 && age <=18 {
		fmt.Println("Teen")
	}else{
		fmt.Println("Child")
	}

	// thing to remember is Go does not have ternary operator yet
}