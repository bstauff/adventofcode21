package day1

import (
	"bufio"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumberOfIncreasesFindsIncreases(t *testing.T) {
	sampleData := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	countOfIncreases := GetNumberOfIncreases(sampleData)

	assert.Equal(t, 7, countOfIncreases)
}

func TestGetNumberOfIncreasesSlidingWindowFindsIncreases(t *testing.T) {
	sampleData := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	countOfIncreases, err := GetNumberOfIncreasesSlidingWindow(sampleData)

	assert.Nil(t, err)

	assert.Equal(t, 5, countOfIncreases)
}

func TestGetNumberOfIncreasesSlidingWindowReturnsErrorIfNoWindows(t *testing.T) {
	sampleData := []int{199, 200}

	countOfIncreases, err := GetNumberOfIncreasesSlidingWindow(sampleData)

	assert.NotNil(t, err)
	assert.Zero(t, countOfIncreases)
}

func TestGetNumberOfIncreasesSlidingWindowWithLargeInput(t *testing.T) {
	file, err := os.Open("input.txt")

	if err != nil {
		t.Errorf("could not read input.txt:%s", err.Error())
	}
	testData :=	make([]int, 0) 
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data, err := strconv.Atoi(scanner.Text())
		if err != nil {
			t.Errorf("failed to parse int to string: %s", err.Error())
		}
		testData = append(testData, data)
	}

	countOfIncreases, err := GetNumberOfIncreasesSlidingWindow(testData)

	assert.Nil(t, err)
	assert.NotZero(t, countOfIncreases)
	t.Logf("input.txt contains %d increasess", countOfIncreases)
}
