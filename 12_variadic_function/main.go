package main

import "fmt"


func sum(nums ...int)int{
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(1,2,3,true,"hey")
	nums := []int{1, 2, 3, 4, 5}
	result1 := sum(nums...)
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of slice:", result1)
	fmt.Println("Sum:", result)
}