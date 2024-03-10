package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductExceptSelf_success(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	result := productExceptSelf(nums)

	assert.Equal(t, expected, result)
}

func TestProductExceptSelf_success_2(t *testing.T) {
	nums := []int{-1, 1, 0, -3, 3}
	expected := []int{0, 0, 9, 0, 0}

	result := productExceptSelf(nums)

	assert.Equal(t, expected, result)
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	result := make([]int, n)

	left[0] = 1
	right[len(right)-1] = 1

	for i := 1; i < n; i++ {
		left[i] = nums[i-1] * left[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	for i := 0; i < n; i++ {
		result[i] = left[i] * right[i]
	}

	return result
}
