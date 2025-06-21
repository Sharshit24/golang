package main

import "fmt"


// iterating over data structures in Go
func main() {
	nums :=[]int{1, 2, 3, 4, 5}


	// for loop

	// for i:=0; i<len(nums); i++{
	// 	fmt.Println(nums[i])
	// }

	sum :=0
	// for range loop
	for _, num := range nums {
		sum += num
		// fmt.Println(num,i)
	}
	// fmt.Println("Sum:", sum)

	// m :=map[string]string{"name":"john", "age":"30"}
	// for k,v :=range m{
	// 	// fmt.Println("Key:", k, "Value:", v)
	// }

	// unicode
	for i,c :=range "golang"{
		fmt.Println(i,c)
	}
}