package main

import "fmt"

func fizzBuzz(n int) string {
	representation := ""
	i := 1

	if n == 0 {
		return "0"
	}

	for i <= n {
		isFactorOf3 := i%3 == 0
		isFactorOf5 := i%5 == 0

		if isFactorOf3 && isFactorOf5 {
			representation += "fizzbuzz"
		} else if isFactorOf5 {
			representation += "buzz"
		} else if isFactorOf3 {
			representation += "fizz"
		} else {
			representation += fmt.Sprint(i)
		}

		i++
	}

	return representation
}

func main() {
	fmt.Println(fizzBuzz(15))
}
