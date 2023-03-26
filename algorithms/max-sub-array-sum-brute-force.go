package main

import (
	"fmt"
	"math"
)

func maxSubArraySum(arr []int, size int) int {
	var maxSum int = math.MinInt

	for i := 0; i < size; i++ {
		var currSum int = 0
		var j int = i

		for j < size {
			currSum += arr[j]
			if currSum > maxSum {
				maxSum = currSum
			}
			j++
		}
	}

	return maxSum

}

func main() {
	fmt.Println(maxSubArraySum([]int{-3, -4, 5, -1, 2, -4, 6, -1}, 8))
	fmt.Println(maxSubArraySum([]int{-2, 3, -1, 2}, 4))
}
