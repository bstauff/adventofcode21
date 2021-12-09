package day2

import (
	"strconv"
	"strings"
	"testing"

	"github.com/bstauff/adventofcode21/internal/pkg/helpers"
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

func TestRelayOrdersToHelmsmanWithAimWithBadCommand(t *testing.T) {
	sub := NewAimSubmarine()

	ordersForHelm := []HelmOrder{
		MakeHelmOrder("bananas", 5),
	}

	err := sub.RelayOrdersToHelmsman(ordersForHelm)

	assert.NotNil(t, err)
}

func TestRelayOrdersToHelmsmanWithAimWithLargeInput(t *testing.T) {

	transformFunc := func(s string) (interface{}, error) {
		split := strings.Split(s, " ")
		command := split[0]
		units, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		return MakeHelmOrder(command, units), nil
	}

	input, err := helpers.ReadInput("input.txt", transformFunc)

	if err != nil {
		t.Error(err.Error())
	}

	orders := make([]HelmOrder, 0, len(input))

	for _, inputData := range input {
		order, ok := inputData.(HelmOrder)
		if !ok {
			t.Error("failed to get HelmOrder from input")
		}
		orders = append(orders, order)
	}

	sub := NewAimSubmarine()
	sub.RelayOrdersToHelmsman(orders)

	result := sub.CurrentPosition()
	assert.Equal(t, 1956047400, result)
}
