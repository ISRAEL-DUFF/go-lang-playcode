package main

import (
	"fmt"

	"sync"
	"time"
)

func setTimeout(fn func(), seconds int) *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(seconds) * time.Second)
		fn()
	}()

	return &wg
}

type IntervalStruct struct {
	wg         *sync.WaitGroup
	shouldStop *bool
}

func (b *IntervalStruct) stopInterval() {
	*(b.shouldStop) = true
}

func setInterval(f func(), seconds int) IntervalStruct {
	var wg sync.WaitGroup
	shouldStop := false

	wg.Add(1)

	t := time.NewTimer(time.Duration(seconds) * time.Second)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-t.C:
				t.Reset(time.Duration(seconds) * time.Second)
				f()
				break
			default:

			}

			if shouldStop {
				break
			}
		}
	}()

	data := IntervalStruct{}

	data.wg = &wg
	data.shouldStop = &shouldStop

	return data
}

func main() {
	fmt.Println("Step 1")
	var tt IntervalStruct
	checkValue := 0

	t1 := setTimeout(func() {
		fmt.Println("1st Timer")
		var i int = 0

		for i < 4 {
			fmt.Println(i)
			i += 1
		}
		checkValue += 1
	}, 10)

	tt = setInterval(func() {
		fmt.Println("Set Interval: Running...")
		if checkValue == 2 {
			fmt.Println("Stopping timer")
			tt.stopInterval()
		}
	}, 2)

	t2 := setTimeout(func() {
		fmt.Println("2nd Timer")
		var i int = 770

		for i < 777 {
			fmt.Println(i)
			i += 1
		}

		checkValue += 1
	}, 12)

	defer t1.Wait()
	defer tt.wg.Wait()
	defer t2.Wait()

	fmt.Printf("Timer Index 1\n")
	fmt.Printf("Timer Index 2 \n")

	// this is neccessary: this main function needs to live for the livespan of those timers
	// time.Sleep(20 * time.Second)

}

// setInterval(func() {
// 	fmt.Println("3rd Timer")
// }, 4)

// time.AfterFunc(5*time.Duration(3), func() {
// 	fmt.Println("Exected funtion here")
// 	var i int = 770

// 	for i < 777 {
// 		fmt.Println(i)
// 		i += 1
// 	}
// })

// time.AfterFunc(5*time.Duration(3), func() {
// 	fmt.Println("Exected funtion here")
// 	var i int = 500

// 	for i < 506 {
// 		fmt.Println(i)
// 		i += 1
// 	}
// })
