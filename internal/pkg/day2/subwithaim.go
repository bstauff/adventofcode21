package day2

type AimSubmarine struct {
	*Submarine
	aim int
}

func NewAimSubmarine() *AimSubmarine {
	sub := new(AimSubmarine)
	sub.Submarine = new(Submarine)
	sub.aim = 0
	return sub
}
