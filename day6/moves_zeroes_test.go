package day6

import (
	"fmt"
	"testing"
)

func TestMovesZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	result := moveZeroes(nums)

	fmt.Println(result)
}

func moveZeroes(nums []int) []int {
	index := 0

	for j, num := range nums {

		if num != 0 {
			nums[index], nums[j] = nums[j], nums[index]
			index++
		}

	}

	return nums
}
