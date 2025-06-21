package main

import (
	"fmt"
	"slices"
)

// slices -> dynamic arrays
// most used construct in Go
// useful methods: append, copy, len, cap
func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println(len(nums))

	// initialize slice with make(not nil)
	nums = make([]int, 5) // 5 is the initial length==capacity
	fmt.Println(len(nums))
	fmt.Println(nums)

	// capacity
	fmt.Println(cap(nums))

	// append to slice
	nums = append(nums, 1)
	fmt.Println(nums)
	fmt.Println(cap(nums))

	// copy function
	nums = make([]int, 0, 5)
	nums = append(nums, 2)
	var nums2 = make ([]int,len(nums) )
	// copy 
	fmt.Println(nums, nums2)
	copy(nums2, nums)
	fmt.Println(nums, nums2)


	// slice operator
	var nums3=[]int{1, 2, 3, 4, 5}
	fmt.Println(nums3[1:3])

	// slice package
	nums4 := []int{1, 2, 3, 4, 5}
	nums5 := []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Equal(nums4, nums5)) // true
}