package main

import "fmt"

// you should always assign values to const when you declare it
const pi = 3.141
func main()  {

	fmt.Println(pi)

	// const grouping declaration
	const (
		host = "local"
		port = 55551
	)
	
	fmt.Println(host)
	fmt.Println(port)
}