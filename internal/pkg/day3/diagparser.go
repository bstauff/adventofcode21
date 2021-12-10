package day3

func GetPowerConsumption(diagReport []uint16) uint64 {
	masks := []uint16{
		0x800,
		0x400,
		0x200,
		0x100,
		0x80,
		0x40,
		0x20,
		0x10,
		0x8,
		0x4,
		0x2,
		0x1,
	}

	numberOfEntries := len(diagReport)
	bitCounts := make([]int, 12)

	for _, report := range diagReport {
		for i, mask := range masks {
			if report&mask == mask {
				bitCounts[i]++
			}
		}
	}

	gamma := uint16(0)

	for i, countOfOnes := range bitCounts {
		if countOfOnes > numberOfEntries-countOfOnes {
			gamma = gamma | masks[i]
		}
	}

	epsilon := uint16(^gamma)
	epsilon = epsilon & 0b0000111111111111

	powerRating := uint64(gamma) * uint64(epsilon)

	return powerRating
}
