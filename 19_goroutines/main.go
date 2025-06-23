package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(id int,w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id) 
}


func main(){
	// waitgroups
	var wg sync.WaitGroup
	for i :=0; i <=10; i++ {
		// go func(i int ) {
		// 	fmt.Println(i)
		// }(i)
		wg.Add(1) // increment the waitgroup counter
		go task(i, &wg)
	}
	wg.Wait() // wait for all goroutines to finish
	fmt.Println("all tasks completed")
	//time.Sleep(2 * time.Second)
}