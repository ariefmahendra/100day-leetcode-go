package day4

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestIncreasingTripletSubsequence(t *testing.T) {
	nums := []int{2, 1, 5, 0, 4, 6}
	result := increasingTriplet(nums)

	assert.True(t, result)
}

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}

	return false
}
