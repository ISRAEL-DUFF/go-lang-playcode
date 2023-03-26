package main

import (
	"fmt"
	"time"
)

func main() {
	myChan := make(chan int, 10) // create a buffered channel of type int and size 10

	go func() {
		for i := 1; i <= 10; i++ {
			myChan <- i // Send the integer i through the channel

			// Do other stuffs
			fmt.Printf("Sent %d\n", i)

		}
		close(myChan) // Close the channel when done

	}()

	fmt.Printf("Received %d\n", <-myChan)

	time.Sleep(time.Second) // this is merely to make the output display slowly.
}
