package main

import (
	"fmt"
	"os"
)

func forLoopsum(i, j int) int {
	var sum = 0
	for k := i; k <= j; k++ {
		sum += k
	}
	return sum
}

func whileLoopsum(i, j int) int {
	var sum = 0

	// A for-loop can be a while-loop
	var k = i
	for k <= j {
		sum += k
		k++
	}

	return sum
}

func echo() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo_two() string {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s = sep + arg
		fmt.Println("Index ", index)
		sep = ""
	}
	return s
}

func main() {
	fmt.Println(forLoopsum(4, 100))
	fmt.Println(whileLoopsum(4, 100))
	fmt.Print(echo())
	fmt.Println(echo_two())

	fmt.Println(os.Args[1:])
}
