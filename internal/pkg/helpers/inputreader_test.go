package helpers

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInputWithGoodData(t *testing.T) {
	transformFunc := func(s string) (interface{}, error) {
		n := strings.Split(s, " ")[1]
		number, err := strconv.Atoi(n)
		if err != nil {
			return 0, err
		}
		return number, nil
	}

	data, err := ReadInput("input.txt", transformFunc)
	if err != nil {
		t.Error(err.Error())
	}

	readNumbers := make([]int, 0, len(data))
	for _, entry := range data {
		s, ok := entry.(int)
		if !ok {
			t.Error("failed to parse read data to number")
		}
		readNumbers = append(readNumbers, s)
	}

	expectedNumbers := []int{8, 1, 9, 8, 5}

	for i, expectedNumber := range expectedNumbers {
		assert.Equal(t, expectedNumber, readNumbers[i])
	}

}
