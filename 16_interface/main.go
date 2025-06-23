package main

import "fmt"


// applying interface
type paymenter interface {
	pay(amount float32)
}
type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw :=razorpay{}
	//stripePaymentGw := stripe{}
	//razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Payment made using Razorpay:", amount)
}


type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Payment made using Stripe:", amount)
}

type fakePayment struct{}
func (f fakePayment) pay(amount float32) {
	fmt.Println("Fake payment made:", amount)
}

func main() {
	//stripePaymentGw := stripe{}
	newPayment := payment{
	gateway: stripe{},
	}
	newPayment.makePayment(100)
}