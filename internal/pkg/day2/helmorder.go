package day2

type HelmOrder struct {
	Command string
	Units   int
}

func MakeHelmOrder(command string, units int) HelmOrder {
	return HelmOrder{
		command,
		units,
	}
}
