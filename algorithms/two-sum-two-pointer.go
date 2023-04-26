package main

import "fmt"

func twoSum(arr []int, target int) (int, int) {
	i := 0
	j := len(arr) - 1

	for i < j {
		currSum := arr[i] + arr[j]

		if currSum == target {
			return i + 1, j + 1
		} else if currSum < target {
			i++
		} else if currSum > target {
			j--
		}
	}

	return i, j
}

func main() {
	fmt.Println(twoSum([]int{1, 3, 4, 8, 9}, 20))
}
