package main

import "fmt"

// we dont have enums is go, in order to create a enum need to use custom types along with const

type OrderStatus int

const(
	Received OrderStatus = iota
	confirmed 
	prepared
	Paid
	Delivered
)

// func changeOrderStatus(status string) {

// 	fmt.Println("changing order status to", status)

// }

func changeOrderStatus(status OrderStatus) {

	fmt.Println("changing order status to", status)

}

func main() {

	changeOrderStatus(Received)
	changeOrderStatus(prepared)
	changeOrderStatus(Delivered)

}