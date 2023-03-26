package main

import "fmt"

// a function that returns more than one ouptut
// it essentially returns two strings
func swap(str1, str2 string) (string, string) {
	return str1 + str2, str1
}

// it can also be writen as:
func swap_two(str1 string, str2 string) (string, string) {
	return str1 + str2, str1
}

// you can also name the return values - as follows
func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x

	// this return is called a "naked" return
	return
}

func main() {
	output1, output2 := swap("Hello", "Golang")
	output11, output22 := swap_two("Golang", "hello")

	fmt.Println(output1, output2)
	fmt.Println(output11, output22)
	fmt.Println(split(4))
}
