package main

import (
	"fmt"
	//"time"
	//"math/rand"
	//"time"
	//"time"
)

//send data to a channel

// func processNum(numChan chan int){
// 	for num := range numChan{
// 		fmt.Println("Processing number:", num)
// 		time.Sleep(1 * time.Second) // Simulate some processing time
// 	}
// 	// fmt.Println("Processing numbers...",<- numChan)
// }

//receive data from a channel

// func sum(result chan int, num1 int,num2 int){
// 	numResult := num1 + num2
// 	result <- numResult // Send the result back to the channel
// }

//goroutine synchronization

// func task(done chan bool){
// 	defer func(){done <- true}()
// 	fmt.Println("Processing task...")
// }


// func emailSender(emailChan chan string, done chan bool) {
// 	for email := range emailChan {
// 		fmt.Println("Sending email to:", email)
// 		// Simulate sending email
// 		time.Sleep(time.Second)
// 	}
// }


func main() {
	//6

	chan1 :=make(chan int)
	chan2 :=make(chan string)

	go func(){
	chan1<-40
}()

go func(){
	chan2 <- "pong"
}()


	for i:=0; i < 2; i++ {
		select{
			case chan1Value := <-chan1:
				fmt.Println("Received from chan1:", chan1Value)
			case chan2Value := <-chan2:	
				fmt.Println("Received from chan2:", chan2Value)
		}
	}


	//5

	// emailChan := make (chan string,100)
	// done := make(chan bool)
	// go emailSender(emailChan,done)
	// for i := 0; i < 10; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com",i)
	// }

	// fmt.Println("Starting email sender...")
	// close(emailChan)  // this is important either get the deadlock
	// <-done

	// emailChan <- "1@example.com"
	// emailChan <- "1@example.com"

	// fmt.Println((<-emailChan))
	// fmt.Println((<-emailChan))   // non blocking

	//<-done   //blocking main goroutine until the emailSender is done


	//4


	// done :=make(chan bool)
	// go task(done) 
	// <-done   // block until the task is done


	//3

	// result :=make (chan int)
	// go sum(result, 5, 10) // Start the sum function in a goroutine
	// res := <-result // Receive the result from the channel
	// fmt.Println("Sum:", res) // Print the result


	//2

	// numChan := make(chan int) // Create a channel for integers
	// go processNum(numChan)
	// for{
	// 	numChan<-rand.Intn(100)
	// }
	//numChan <- 42 // Send a number to the channel
	//time.Sleep(2 * time.Second) // Wait for a second to ensure the goroutine processes the number


	//1

	//messageChan := make(chan string)
	//messageChan <- "Hello, World!" // Send a data to the channel

	//msg := <-messageChan           // Receive data from the channel

	//fmt.Println(msg) // Print the received message
}