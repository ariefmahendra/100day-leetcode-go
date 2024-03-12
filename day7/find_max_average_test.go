package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxAverage_success(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	expected := 12.75000

	result := findMaxAverage(nums, k)

	assert.Equal(t, expected, result)
}

func findMaxAverage(nums []int, k int) float64 {
	if k > len(nums) {
		return -1.0
	}

	maxSum := 0
	tempSum := 0

	// init sum
	for i := 0; i < k; i++ {
		tempSum += nums[i]
	}

	//init start Index and last Index
	startIndex := 0
	lastIndex := k

	for lastIndex < len(nums) {
		// remove left num
		tempSum -= nums[startIndex]
		startIndex++

		// + right num
		tempSum += nums[lastIndex]
		lastIndex++

		if tempSum > maxSum {
			maxSum = tempSum
		}
	}

	return float64(maxSum) / float64(k)
}
