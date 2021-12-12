package day3

import (
	"strconv"
	"testing"

	"github.com/bstauff/adventofcode21/internal/pkg/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGetPowerConsumptionHasGoodResultsWithLargeInput(t *testing.T) {
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
	powerConsumption := GetPowerConsumption(castInput, 0b0000111111111111)

	assert.Equal(t, uint64(2261546), powerConsumption)
}

func TestGetPowerConsumptionHasGoodResults(t *testing.T) {
	input := []uint16 {
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}

	powerConsumption := GetPowerConsumption(input, 0b0000000000011111)

	assert.Equal(t, uint64(198), powerConsumption)
}

func TestGetLifeSupportRatingHasGoodResults(t *testing.T) {
	input := []uint16 {
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}

	lifeSupportRating := GetLifeSupportRating(input, 0b0000000000011111)

	assert.Equal(t, uint64(230), lifeSupportRating)
}
