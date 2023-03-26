package main

import (
	"fmt"
	"math"
)

func removeDuplicates(nums []int) int {
	var i int = 0
	var currPosition int = 0
	var size int = len(nums)
	var currNum int = math.MinInt

	if size == 0 {
		return 0
	}

	for i < size {
		if currNum != nums[i] {
			nums[currPosition] = nums[i]
			currPosition += 1
			currNum = nums[i]
		}

		i++
	}

	return currPosition
}

func main() {
	var arr []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(arr)
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)

}
