package day5

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCompress_success(t *testing.T) {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	expected := 6

	result := compress(chars)

	assert.Equal(t, expected, result)
}

func compress(chars []byte) int {
	readIndex, writeIndex := 0, 0
	for readIndex < len(chars) {
		currentChar := chars[readIndex]
		currentCount := 0

		for readIndex < len(chars) && currentChar == chars[readIndex] {
			readIndex++
			currentCount++
		}

		chars[writeIndex] = currentChar
		writeIndex++

		if currentCount > 1 {
			count := strconv.Itoa(currentCount)
			for _, i := range count {
				chars[writeIndex] = byte(i)
				writeIndex++
			}
		}
	}

	return writeIndex
}
