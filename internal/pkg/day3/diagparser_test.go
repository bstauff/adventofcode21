package day3

import (
	"strconv"
	"testing"

	"github.com/bstauff/adventofcode21/internal/pkg/helpers"
	"github.com/stretchr/testify/assert"
)

func TestParseHasGoodResults(t *testing.T) {
	transformFunc := func(s string) (interface{}, error) {
		parsedBytes, err := strconv.ParseUint(s, 2, 16)

		if err != nil {
			return nil, err
		}
		return parsedBytes, nil
	}

	input, err := helpers.ReadInput("input.txt", transformFunc)
	if err != nil {
		t.Error(err.Error())
	}

	castInput := make([]uint16, len(input))
	for i, data := range input {
		castInput[i] = uint16(data.(uint64))
	}
	powerConsumption := GetPowerConsumption(castInput)

	assert.Equal(t, uint64(2261546), powerConsumption)
}
