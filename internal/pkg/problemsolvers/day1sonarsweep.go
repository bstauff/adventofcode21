package problemsolvers

import (
	"errors"
	"strconv"
	"strings"
)

type Day1SonarSweep struct {
	depthMeasurements []int
}

func (sweeper *Day1SonarSweep) Initialize(initData string) error {
	measurements, parseErr := parseCommaSeparatedStringToInts(initData)

	if parseErr != nil {
		return errors.New("failed to initialize: " + parseErr.Error())
	}

	sweeper.depthMeasurements = measurements

	return nil
}

func parseCommaSeparatedStringToInts(source string) ([]int, error) {
	splitSource := strings.Split(source, ",")
	parsedData := make([]int, 0, len(splitSource))
	for _, dataString := range splitSource {
		intToAppend, convertErr := strconv.Atoi(dataString)
		if convertErr != nil {
			return nil, convertErr
		}

		parsedData = append(parsedData, intToAppend)
	}
	return parsedData, nil
}

func (sweeper Day1SonarSweep) Solve() (string, error) {
	if sweeper.depthMeasurements == nil {
		return "", errors.New("problemsolver was not initialized with data")
	}

	increases := sweeper.getNumberOfIncreases()

	return strconv.Itoa(increases), nil
}

func (sweeper Day1SonarSweep) getNumberOfIncreases() int {
	lastMeasurement := sweeper.depthMeasurements[0]
	numberOfIncreases := 0

	for _, value := range sweeper.depthMeasurements {
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
