package main

import (
	"fmt"
)

func main() {
	myChan := make(chan int) // create a channel of type int

	go func() {
		for i := 1; i <= 10; i++ {
			myChan <- i // Send the integer i through the channel

			// Do other stuffs
			fmt.Printf("Sent %d\n", i)

		}
		close(myChan) // Close the channel when done

	}()

	fmt.Printf("Received %d\n", <-myChan)
}
