package main

import (
	"fmt"
	"time"
)

// Step 1: Initilaise a channel by make method
//....done := make(chan <- bool, 2) - You can only send values to your channel
//....done := make(<- chan  bool, 2) - You can only receive values to your channel

// step 2: set the Channel to true/ 1 on the completion of routine.
//....done <- true

// step 3: Notify the completion
//...<-done

func main() {
	start := time.Now()
	done := make(chan bool, 2) // Step 1: Initilaise a channel by make method

	//done := make(chan <- bool, 2) - You can only send values to your channel

	//done := make(<- chan  bool, 2) - You can only receive values to your channel

	go func() {
		fmt.Println("Go Routine :Operational task - 1")
		done <- true // step 2: set the Channel to true/ 1 on the completion of routine.

	}()
	<-done // step 3: Notify the completion

	go func() {
		fmt.Println("Go Routine : Operational task - 2")
		done <- true
	}()

	fmt.Println("operational task - 3")
	<-done

	fmt.Println("time taken for execution: ", time.Since(start))
	fmt.Println("Database operation starts")
}
