package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidParentheses_success(t *testing.T) {

	str := "{{}}[]"
	parentheses := validParentheses(str)

	assert.True(t, parentheses)
}

func validParentheses(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var stack []rune
	bracketMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, r := range s {

		_, ok := bracketMap[r]

		if ok {
			stack = append(stack, r)
			continue
		}

		l := len(stack)

		if len(stack) > 0 && bracketMap[stack[l-1]] == r {
			stack = stack[:l-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
