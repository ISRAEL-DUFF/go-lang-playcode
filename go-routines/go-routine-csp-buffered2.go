package main

import (
	"math/rand"
	"time"
)

type Counter struct {
	scores map[int]int
}

type Score struct {
	id    int
	score int
}

type ScoreRequest struct {
	response chan []Score
}

func main() {
	scoreChan := make(chan Score)

	// generate random scores and id and send them to scoreChan every 1ms
	for i := 0; i < 100; i++ {
		go func() {
			id := rand.Intn(3)
			score := rand.Intn(100)

			time.Sleep(5 * time.Millisecond)

			scoreChan <- Score{id, score}
		}()
	}

}
