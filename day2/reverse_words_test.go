package day2

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReverseWords_success(t *testing.T) {
	str := "  hello  world  "
	expected := "world hello"
	result := reverseWords(str)

	assert.Equal(t, expected, result)
}

func reverseWords(s string) string {
	sSlice := strings.Fields(s)

	left, right := 0, len(sSlice)-1

	for left < right {
		sSlice[left], sSlice[right] = sSlice[right], sSlice[left]
		left++
		right--
	}

	return strings.Join(sSlice, " ")
}
