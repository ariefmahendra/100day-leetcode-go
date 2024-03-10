package day3

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMergeAlternately_success(t *testing.T) {
	word1 := "ab"
	word2 := "pqrs"

	expected := "apbqrs"

	result := mergeAlternately(word1, word2)

	assert.Equal(t, expected, result)
}

func mergeAlternately(word1 string, word2 string) string {
	word1Split := strings.Split(word1, "")
	word2Split := strings.Split(word2, "")

	var result []string

	ptr1, ptr2 := 0, 0
	for ptr1 < len(word1Split) && ptr2 < len(word2Split) {
		result = append(result, word1Split[ptr1])
		result = append(result, word2Split[ptr2])
		ptr1++
		ptr2++
	}

	if len(word1Split) > len(word2Split) {
		result = append(result, word1Split[len(word2Split):]...)
	}

	if len(word2Split) > len(word1Split) {
		result = append(result, word2Split[len(word1Split):]...)
	}

	return strings.Join(result, "")
}
