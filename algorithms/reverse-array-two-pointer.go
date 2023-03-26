package main

import (
	"fmt"
)

func swap(arr []string, i int, j int) {
	var temp string = arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func reverse(arr []string, size int) {
	var i int = 0
	var j int = size - 1

	for i < j {
		swap(arr, i, j)

		i++
		j--
	}
}

func main() {
	var arr []string = []string{"a", "b", "c", "d", "e"}

	fmt.Println(arr)
	reverse(arr, 5)
	fmt.Println(arr)
}
