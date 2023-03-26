package main

import "fmt"

func refToLocalVariable() *int {
	var y int = 20
	fmt.Println("Y = ", y)
	return &y
}

func modifyPtrVariable(ptr *int) {
	fmt.Println("Old value: ", *ptr)
	*ptr = 10
	fmt.Println("New value: ", *ptr)
}

func usingNewToCreatePtr() *int {
	var x = new(int)
	*x = 17
	fmt.Println("X = ", *x)
	return x
}

func main() {
	fmt.Println("Hello Pointers")
	var x int = 5
	var p *int = &x

	fmt.Println(*p)
	fmt.Println(x)

	*p = 16 // same as x = 16

	fmt.Println(x)

	// get address of local variable
	var z *int = refToLocalVariable()
	*z = 40
	fmt.Println("Y = ", *z)
	modifyPtrVariable(z)

	var t = usingNewToCreatePtr()
	*t = 60
	fmt.Println("new X = ", *t)
}
