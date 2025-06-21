package main

import "fmt"

func main() {
	// var nums [5]int
	// fmt.Println(len(nums))
	// int->0 ,string->"", bool->false 

	// int
	var nums [5]int
	nums[0] = 1
	fmt.Println(nums)

	// bool
	var nums2[5]bool
	nums2[2]=true
	fmt.Println(nums2)

	// string
	var names [5]string
	names[0] = "John"
	fmt.Println(names)

	// to declare and initialize an array
	nums3 := [5]int{0, 1, 2,3, 4}
	fmt.Println(nums3)


	// 2D array
	nums4 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums4)

	// fixed size, that is predictable
	// memory optimization
	// constant time access
}