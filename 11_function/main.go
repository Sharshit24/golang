package main

import "fmt"

// function
func add(a int,b int)int{
	return a + b
}


func main(){
	nums := []int{4, 5}
	s := add(nums[0], nums[1])
	r := add(4,5)
	fmt.Println(r,s)

}