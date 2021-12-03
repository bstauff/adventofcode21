package problemsolvers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThrowsErrorIfNotInitialized(t *testing.T) {
	solver := new(Day1SonarSweep)

	_, err := solver.Solve()

	assert.NotNil(t, err)
}
func TestThrowsErrorIfGivenBadInitData(t *testing.T) {

	solver := new(Day1SonarSweep)

	initErr := solver.Initialize("apples")

	assert.NotNil(t, initErr)
}
func TestCountIncreasesFindsIncreases(t *testing.T) {
	solver := new(Day1SonarSweep)
	sampleData := "199,200,208,210,200,207,240,269,260,263"
	initError := solver.Initialize(sampleData)

	assert.Nil(t, initError)

	countOfIncreases, solveError := solver.Solve()

	assert.Nil(t, solveError)

	assert.Equal(t, "7", countOfIncreases)
}
