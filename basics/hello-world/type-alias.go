package main

import "fmt"

// type Alias begins with the following syntax
// type [new-name] [existing-type]
type WholeNumber int // i.e wholeNumbers is now same as int
type NaturalNumber int

func main() {
	fmt.Println("Hello Type Alias")
	var x WholeNumber = 30
	var n NaturalNumber = 25
	fmt.Println("X = ", x)
	fmt.Println("n = ", n)

	// NOTE: if you compare x and n, you get a compile error:
	// if x == n ... error
}
