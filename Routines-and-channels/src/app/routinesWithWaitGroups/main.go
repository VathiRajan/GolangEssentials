package main

import (
	"fmt"
	"sync"
)

//Step 1 :- Declare a wait group
//....var wg = sync.WaitGroup{}

// STep 2: Add the no of wait groups
//....wg.Add(1)

//Step 3: Update the completion in the wait group caling the Done method
//....wg.Done()

//Step 4: Make the Main thread to wait until the routine completion
//wg.Wait()

var wg = sync.WaitGroup{} //Step 1 :- Declare a wait group
//Wait group is added globally here since its shared throughout the application

func main() {
	msg := "Hello"

	wg.Add(1) // STep 2: Add the no of wait groups

	go func(bingo string) {

		fmt.Println(bingo)
		wg.Done() //Update the completion in the wait group caling the Done method
	}(msg)

	msg = "Good - Bye"
	wg.Wait() //Make the Main thread to wait until the routine completion
}
