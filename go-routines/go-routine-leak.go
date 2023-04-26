package main

import (
	"fmt"
)

// EXAMPLE GO ROUTINE LEAK
// How did the leak happen?
// ANS: the main function passed in a 'nil' to the do work function, Hence the 'strings' channel will never actually get any strings written to it
//		So that the go-routin containing doWork will remain in memory for the entire life-time of the program
// How do we know there is a leak?
// ANS: You will notice in the output that the string "Leaked goroutin" didn't print, hence we never got the the end of the goroutine.

// NOTE: we can easily get into a deadlock by simply joinning doWork back into main.... See example below

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completedChan := make(chan interface{})

		go func() {
			defer fmt.Println("Leaked goroutin")
			defer close(completedChan)

			// use the read from the channel you have received
			for s := range strings {
				// do something with the read string
				fmt.Println(s)
			}
		}()

		return completedChan
	}

	doWork(nil)

	// more work can be done here

	fmt.Println("Done.")

}
