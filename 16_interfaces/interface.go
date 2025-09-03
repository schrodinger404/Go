
package main

import "fmt"

type paymenter interface{
	pay(amount float32)
}

// all of this is to understand strcts and their methods, 
// if you dont use ineterface you'll alway be changing the gateway in the in the payment struct, use paymenter unterface as defined above
type paymet struct{

	//gateway stripe
	//gateway razorpay
	// instead of hardcoding the gateway here, you can use the interface
	gateway paymenter
}

func (p paymet) makePayment(amount float32) {

	// razorPayGw := razorpay{}

	// razorPayGw.pay(amount)
	p.gateway.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic for razor pay 
	fmt.Println("making payment using razor pay", amount)

}

type stripe struct{}

func (s stripe) pay(amount float32)  {
	// logic for strip pay
	fmt.Println("making payment using stripe", amount)

	
}

func main() {

	// now you can initialize structs
	stripeGw := stripe{}
	//razorpayGw := razorpay{}


	// payment struct implementation
	Mypay := paymet{
		gateway: stripeGw,
	}

	Mypay.makePayment(2345)
}
