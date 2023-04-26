// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import (
	"fmt"
	"time"
)

/*
	HOW DO WE SOLVE GO-ROUTINE LEAK?

	ANS: The way to successfully mitigate go-routine leak is to establish a signal between the parent gorou‐tine and its children
	  in such a way that allows the parent to signal cancellation to its children.
	  By convention, this signal is usually a read-only channel named 'done'.
	  The parent gorou‐tine passes this channel to the child goroutine and then closes the channel when it wants to cancel the child goroutine.

	  Example is given below:
*/

func main() {
	// NOTE: by convention, the done channel is always the 1st parameter to a function as shown below

	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited!!!")
			defer close(terminated)

			for {
				select {
				case <-done:
					return
				case s := <-strings:
					// do something interesting with channel strings
					fmt.Println(s)

				}
			}
		}()

		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	// let's cancel the goroutine after 2secs
	go func() {
		time.Sleep(time.Second * 2)
		print("Canceling dowork go-routine...")

		// don't forget to close the done channel
		close(done)
	}()

	// join the goroutine spawned from doWork above with this main gorou‐tine
	<-terminated

	fmt.Println("Done.")
}
