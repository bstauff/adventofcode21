package day3

func GetPowerConsumption(diagReport []uint16, resultBitMask uint16) uint64 {
	gamma := getMostCommonBits(diagReport)
	epsilon := uint16(^gamma)
	epsilon = epsilon & resultBitMask

	powerRating := uint64(gamma) * uint64(epsilon)

	return powerRating
}

func GetLifeSupportRating(diagReport []uint16, resultBitMask uint16) uint64 {
	mostCommonBits := getMostCommonBits(diagReport)

	masks := []uint16{
		// 0x800,
		// 0x400,
		// 0x200,
		// 0x100,
		// 0x80,
		// 0x40,
		// 0x20,
		0x10,
		0x8,
		0x4,
		0x2,
		0x1,
	}
	trimmed := make([]uint16, 0, len(diagReport))
	copy(trimmed, diagReport)
	for _, mask := range masks {
		
	}



	oxygenRating := recursiveFilter(mostCommonBits, masks, 0, diagReport)

	return uint64(oxygenRating[0])
}

func recursiveFilter(commonBits uint16, masks []uint16, maskIndex int, diagReports []uint16) []uint16 {
	if len(diagReports) == 1 || maskIndex == len(masks) {
		return diagReports
	}

	trimmed := filterWithMask(masks[maskIndex], diagReports, commonBits)

	return recursiveFilter(commonBits, masks, maskIndex + 1, trimmed)
}




func filterWithMask(bitmask uint16, diagReports []uint16, commonBits uint16) []uint16 {
	trimmedReports := make([]uint16, 0, len(diagReports))

	for _, report := range diagReports {
		bitvalue := report & bitmask

		shouldGrabReport := bitvalue^(commonBits&bitmask) == 0

		if shouldGrabReport {
			trimmedReports = append(trimmedReports, report)
		}
	}
	return trimmedReports
}

func getMostCommonBits(diagReport []uint16) uint16 {
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

	mostCommonBits := uint16(0)

	for i, countOfOnes := range bitCounts {
		if countOfOnes > numberOfEntries-countOfOnes {
			mostCommonBits = mostCommonBits | masks[i]
		}
	}

	return mostCommonBits
}
