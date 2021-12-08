package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDriveSubGiveGoodPosition(t *testing.T) {
	sub := NewSubmarine()

	sub.MoveForward(5)
	sub.Descend(5)
	sub.MoveForward(5)
	sub.Ascend(3)
	sub.Descend(8)
	sub.MoveForward(2)

	expectedPosition := 150

	assert.Equal(expectedPosition, sub.CurrentPosition())
}