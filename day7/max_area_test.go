package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea_success(t *testing.T) {
	container := []int{1, 1}
	expected := 1

	result := maxArea(container)

	assert.Equal(t, expected, result)
}

func TestMaxArea_success_2(t *testing.T) {
	container := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49

	result := maxArea(container)

	assert.Equal(t, expected, result)
}

func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		tempHeight := 0
		tempWidth := right - left
		if height[left] < height[right] {
			tempHeight = height[left]
			left++
		} else {
			tempHeight = height[right]
			right--
		}

		tempArea := tempWidth * tempHeight

		if tempArea > maxArea {
			maxArea = tempArea
		}
	}

	return maxArea
}
