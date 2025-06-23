package main

import "fmt"

// enumerated types
type OrderStatus int 
const(
	Received OrderStatus = iota // 0
	confirmed                  // 1
	Prepared				  // 2
	Shipped                   // 3
	Delivered                 // 4
)


func changeOrderStatus(status OrderStatus){
	fmt.Println("Order status changed to:", status)
}

func main(){
	changeOrderStatus(Delivered)
}