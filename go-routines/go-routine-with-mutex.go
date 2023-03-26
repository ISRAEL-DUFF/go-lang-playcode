package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	score int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	// acquire a lock
	c.mu.Lock()

	// Do stuff here e.g
	c.score += 10

	// Release lock
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()

	defer c.mu.Unlock()

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
