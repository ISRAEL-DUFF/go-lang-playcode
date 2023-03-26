package main

import "fmt"

// you can use var to declare a list of variables as follows
// notice that the type comes last
// And also, these variables have default values of false
var java, python, php bool

// they can also have initializers
var ruby, haskell, javascript bool = true, false, true

// they don't even have to be the same type
var lisp, perl, dart = "Hello", 3, true

func main() {
	// this has a default value of zero
	var i int

	fmt.Println(i, java, python, php)
	fmt.Println(ruby, haskell, javascript)
	fmt.Println(lisp, perl, dart)

	// this can only be done within functions
	a, b, c := "Hello", true, 20
	d := 40

	fmt.Println(a, b, c)
	fmt.Println(d)
}
