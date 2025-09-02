package main

import (

	"fmt"
)

// struct is  user defined datatype, just like classes in Java
// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.
// OOP can be acieved using structs

// this is simple struct
type Order struct {
	number int
	amount float32
	status string
}


// cunstructor
func NewOrder(number int, amount float32, status string) *Order {
	myOrder := Order{
		number: number,
		amount: amount,
		status: status,
	}	
	return &myOrder
}

// and the fuction associated with struct
func (o *Order) changeStatus(status string) {
	o.status = status

}

func (o Order) getAmount() float32 {
	return o.amount
}

func main() {

	myOrder := Order{
		number: 1,
		amount: 100.0,
		status: "received",
	}

	fmt.Println(myOrder)

	myOrder.changeStatus("sent")
	fmt.Println(myOrder)

	// you can itialize the value you wanted, all not mandatory
	myOrder2 := Order{
		number: 10,
		amount: 200.0,
	}

	fmt.Println(myOrder2)
	fmt.Println(myOrder2.getAmount())

	// creating strcut using constructor

	order3 := NewOrder(33,23.4,"paid")

	fmt.Println(order3)

}
