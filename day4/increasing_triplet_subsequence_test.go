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
	firstNumber, secondNumber := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num <= firstNumber {
			firstNumber = num
		} else if num <= secondNumber {
			secondNumber = num
		} else {
			return true
		}
	}

	return false
}
