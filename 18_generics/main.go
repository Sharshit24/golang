package main

import "fmt"

func printSlice(items []int) {
	for _, item := range items {
		fmt.Println (item)
	}
}
func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println (item)
	}
}

// generics on functions
func printSliceGeneric[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// generics on structs
type stack[T any] struct {
	elements []T
}

func main(){
	printSlice([]int{1, 2, 3, 4, 5})
	names := []string{"Alice", "Bob", "Charlie"}
	printStringSlice(names)
}