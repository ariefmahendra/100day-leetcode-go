package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKidsWithCandies_success(t *testing.T) {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	expected := []bool{true, true, true, false, true}

	result := kidsWithCandies(candies, extraCandies)

	assert.Equal(t, expected, result)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var result []bool

	maxValue := 0
	for _, candy := range candies {
		if candy > maxValue {
			maxValue = candy
		}
	}

	for _, candy := range candies {
		temp := candy + extraCandies
		if temp >= maxValue {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}
