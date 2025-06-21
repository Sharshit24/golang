package main

import "fmt"

// Maps -> key-value pairs
// Maps are unordered collections of key-value pairs
func main() {
	// creating a map
	m := make(map[string]string)

	// adding key-value pairs
	m["name"] = "golang"
	m["area"] = "backend"
	m["price"] = "free"
	fmt.Println(m["name"], m["area"], m["price"])

	// deleting a key-value pair
	delete(m, "price")
	fmt.Println(m["name"], m["area"], m["price"])

	// clearing a map
	clear(m)
	fmt.Println(m["name"], m["area"], m["price"])
 
	// length of map
	// fmt.Println(len(m))

	n := map[string]int{"price":50, "quantity": 100}
	// fmt.Println(n)
	k,ok := n["price"]
	fmt.Println(k)
	if ok {
		fmt.Println("all ok")
}else{
		fmt.Println("not ok")
	}
}
