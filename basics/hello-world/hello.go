package main

import (
	"fmt"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func subtract(x, y, z int) int {
	return add(x, y) - z
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println("My Random number is:", rand.Intn(13))
	fmt.Println("My Random number is:", rand.Intn(13))

	fmt.Println("Addition: ", add(4, -7))
	fmt.Println("Subtraction: ", subtract(1, 2, 3))
}
