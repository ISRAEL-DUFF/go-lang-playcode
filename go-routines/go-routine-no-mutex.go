package main

import (
	"fmt"
	"time"
)

type Counter struct {
	score int
}

func (c *Counter) Inc() {
	c.score += 10
}

func (c *Counter) Value() int {
	return c.score
}

func main() {
	var counter = Counter{}

	// spawn 10 routines to increment the counter
	for i := 1; i <= 10; i++ {
		go counter.Inc()
	}

	time.Sleep(time.Second)
	fmt.Println(counter.Value())
}
