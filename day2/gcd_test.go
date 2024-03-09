package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGdc_success(t *testing.T) {
	str1 := "ABCABCABC"
	str2 := "ABC"

	result := gcdOfStrings(str1, str2)

	assert.Equal(t, "ABC", result)
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	}

	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}

	if str1[:len(str2)] != str2 {
		return ""
	}

	return gcdOfStrings(str1[len(str2):], str2)
}
