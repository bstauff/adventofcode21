package day1

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
