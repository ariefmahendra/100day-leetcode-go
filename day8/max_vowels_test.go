package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxVowels_success(t *testing.T) {
	input := "abciiidefiiee"
	k := 4
	expected := 4
	result := maxVowels(input, k)

	assert.Equal(t, expected, result)
}

func TestMaxVowels_success_2(t *testing.T) {
	input := "leetcode"
	k := 3
	expected := 2
	result := maxVowels(input, k)

	assert.Equal(t, expected, result)
}

func maxVowels(s string, k int) int {
	l, r := 0, k
	maxCount := 0

	vowel := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'u': true,
		'o': true,
	}

	for i := 0; i < k; i++ {
		if vowel[rune(s[i])] {
			maxCount++
		}
	}

	tempCount := maxCount

	for r < len(s) {
		// tambahkan count pada next index
		if vowel[rune(s[r])] {
			tempCount++
		}

		// kurangi count dengan index awal
		if vowel[rune(s[l])] {
			tempCount--
		}

		if tempCount > maxCount {
			maxCount = tempCount
		}

		l++
		r++
	}

	return maxCount
}
