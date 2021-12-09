package day2

import "errors"

type AimSubmarine struct {
	horizontalPostiion int
	depth              int
	aim                int
}

func NewAimSubmarine() *AimSubmarine {
	sub := new(AimSubmarine)
	sub.horizontalPostiion = 0
	sub.depth = 0
	sub.aim = 0
	return sub
}

func (sub *AimSubmarine) moveForward(distance int) {
	sub.horizontalPostiion += distance
	sub.depth += (sub.aim * distance)
}

func (sub *AimSubmarine) descend(depth int) {
	sub.aim += depth
}

func (sub *AimSubmarine) ascend(depth int) {
	sub.aim -= depth
}

func (sub AimSubmarine) CurrentPosition() int {
	return sub.depth * sub.horizontalPostiion
}

func (sub *AimSubmarine) RelayOrdersToHelmsman(orders []HelmOrder) error {
	for _, helmOrder := range orders {
		err := sub.handleOrder(helmOrder)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sub *AimSubmarine) handleOrder(order HelmOrder) error {
	switch order.Command {
	case "forward":
		sub.moveForward(order.Units)
	case "up":
		sub.ascend(order.Units)
	case "down":
		sub.descend(order.Units)
	default:
		return errors.New("received unknown command")
	}
	return nil
}