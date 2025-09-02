package main


// im simply shwing a constructor here, no where used
// simple struct
type Person struct{
	 name string
	 age int
}

// we dont have constructor in Go, so use func to define a cunstructor
func NewPerson(name string, age int) *Person {
	// initial setup goes here
	P := Person{
		name: name,
		age: age,
	}
	return &P
}
