package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumberOfIncreasesFindsIncreases(t *testing.T) {
	sampleData := []int{199,200,208,210,200,207,240,269,260,263}
	countOfIncreases := GetNumberOfIncreases(sampleData)

	assert.Equal(t, 7, countOfIncreases)
}

func TestGetNumberOfIncreasesSlidingWindowFindsIncreases(t *testing.T) {
	sampleData := []int{199,200,208,210,200,207,240,269,260,263}
	countOfIncreases := GetNumberOfIncreasesSlidingWindow(sampleData)

	assert.Equal(t, 5, countOfIncreases)
}
