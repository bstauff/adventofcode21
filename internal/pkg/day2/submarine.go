package day2

import "errors"

type Submarine struct {
	horizontalPostiion int
	depth              int
}

func NewSubmarine() *Submarine {
	sub := new(Submarine)
	sub.horizontalPostiion = 0
	sub.depth = 0

	return sub
}

func (sub *Submarine) moveForward(distance int) {
	sub.horizontalPostiion += distance
}

func (sub *Submarine) descend(depth int) {
	sub.depth += depth
}

func (sub *Submarine) ascend(depth int) {
	sub.depth -= depth
}

func (sub Submarine) CurrentPosition() int {
	return sub.depth * sub.horizontalPostiion
}

func (sub *Submarine) RelayOrdersToHelmsman(orders []HelmOrder) error {
	for _, helmOrder := range orders {
		err := sub.handleOrder(helmOrder)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sub *Submarine) handleOrder(order HelmOrder) error {
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
