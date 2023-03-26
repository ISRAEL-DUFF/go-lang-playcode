package main

import "fmt"

func main() {
	var numbers [5]int
	var numbers2 [5]int = [5]int{4, 1, 6}

	for i, j := range numbers {
		fmt.Println(i, j)
	}

	for _, j := range numbers2 {
		fmt.Println(j)
	}
}
