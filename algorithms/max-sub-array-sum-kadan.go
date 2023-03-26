package main

import (
	"fmt"
	"math"
)

// Kadane’s Algorithm is an iterative dynamic programming algorithm.
// It calculates the maximum sum subarray ending at a particular position by using the maximum sum subarray ending at the previous position

func maxSubArraySum(arr []int, size int) int {
	var maxSum = math.MinInt
	var currSum = 0

	for i := 0; i < size; i++ {
		currSum += arr[i]

		if currSum > maxSum {
			maxSum = currSum
		}

		if currSum < 0 {
			currSum = 0
		}
	}

	return maxSum
}

func main() {
	fmt.Println(maxSubArraySum([]int{-3, -4, 5, -1, 2, -4, 6, -1}, 8))
	fmt.Println(maxSubArraySum([]int{-2, 3, -1, 2}, 4))
}
