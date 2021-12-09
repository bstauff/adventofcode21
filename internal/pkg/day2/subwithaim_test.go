package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelayOrdersToHelmsmanWithAim(t *testing.T) {
	sub := NewAimSubmarine()

	ordersForHelm := []HelmOrder{
		MakeHelmOrder("forward", 5),
		MakeHelmOrder("down", 5),
		MakeHelmOrder("forward", 8),
		MakeHelmOrder("up", 3),
		MakeHelmOrder("down", 8),
		MakeHelmOrder("forward", 2),
	}

	err := sub.RelayOrdersToHelmsman(ordersForHelm)

	assert.Nil(t, err)

	expectedPosition := 900

	assert.Equal(t, expectedPosition, sub.CurrentPosition())
}
