package day7

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestKSumPair(t *testing.T) {

	nums := []int{1, 4, 3, 5, 6, 7, 2, 1}
	k := 5
	expected := 2

	result := maxOperations(nums, k)

	assert.Equal(t, expected, result)
}

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1
	op := 0

	for left < right {
		tempSum := nums[left] + nums[right]

		if tempSum == k {
			op++
			left++
			right--
		} else if tempSum > k {
			right--
		} else if tempSum < k {
			left++
		}
	}

	return op
}
