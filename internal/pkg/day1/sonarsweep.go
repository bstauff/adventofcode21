package day1

import "errors"

func GetNumberOfIncreases(depthMeasurements []int) int {
	lastMeasurement := depthMeasurements[0]
	numberOfIncreases := 0

	for _, value := range depthMeasurements {
		if didMeasurementIncrease(lastMeasurement, value) {
			numberOfIncreases++
		}
		lastMeasurement = value
	}

	return numberOfIncreases
}

func didMeasurementIncrease(previousMeasurement int, newMeasurement int) bool {
	return newMeasurement > previousMeasurement
}

func GetNumberOfIncreasesSlidingWindow(depthMeasurements []int) (int, error) {
	windowCount := len(depthMeasurements) - 2

	if windowCount < 1 {
		return 0, errors.New("depthMeasurements must have at least 1 window")
	}

	increasingWindowsCount := 0

	previousWindowSum := sumWindowFromIndex(depthMeasurements, 0)

	for index := 1; index < windowCount; index++ {
		windowSum := sumWindowFromIndex(depthMeasurements, index)
		if windowSum > previousWindowSum {
			increasingWindowsCount++
		}
		previousWindowSum = windowSum
	}

	return increasingWindowsCount, nil
}

func sumWindowFromIndex(depthMeasurements []int, index int) int {
	return depthMeasurements[index] + depthMeasurements[index+1] + depthMeasurements[index+2]
}
