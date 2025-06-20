package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		// fmt.Println(i)
		i++
	}
	// infinite loop
	// for{
	// 	// fmt.Println("1")
	// }
	for i := 0; i < 3; i++ {
		// break
		if i == 2{
			continue
		}
		fmt.Println(i)
	}
	age := 20
	if age>=18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	role := "admin"
	haveAccess := false
	if role== "admin" || haveAccess {
		fmt.Println("You have access")
	} else {
		fmt.Println("You do not have access")
	}
}