package main

import (
	"fmt"
	"time"
)

//  order struct
type Order struct {
	id 		string
	amount 	float32
	status 	string
	createdAt time.Time  // nanosecond precision
	received bool
}

func newOrder (id string, amount float32,status string)*Order{ 
	myOrder1 := Order{
		id:        id,
		amount:    amount,
		status:    status,
		// createdAt: time.Now(),
		// received:  false,
	}
	return &myOrder1
}

func (o *Order) changeStatus(status string) {
	o.status = status
}

func main() {
	myOrder1 :=newOrder("12345", 99.99, "Pending") // Create a new order using constructor function
	fmt.Println(myOrder1) // Print the order details
	myOrder:= Order{
		id:        "12345",
		amount:    99.99,
		status:    "Pending",
		createdAt: time.Now(),
		received : true,
	}
	myOrder.changeStatus("Shipped") // Change the order status using method


	// myOrder.status = "Shipped" // Update the order status


	// Print the order details
	fmt.Println("Order ID:", myOrder)
}