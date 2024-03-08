package day1

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestLongestCommonPrefix_success(t *testing.T) {
	var strs = []string{"flower", "flow", "flight"}

	result := longestCommonPrefix(strs)

	assert.Equal(t, "fl", result)
}

func TestLongestCommonPrefix_success_2(t *testing.T) {
	var strs = []string{"dog", "racecar", "car"}

	result := longestCommonPrefix(strs)

	assert.Equal(t, "", result)
}

func longestCommonPrefix(words []string) string {

	// check length of words
	if len(words) == 1 {
		return words[0]
	}

	sort.Strings(words)

	//compare first and last
	l := len(words)
	for i := range words[0] {
		if words[0][i] != words[l-1][i] {
			return words[0][:i]
		}
	}

	return words[0]
}
