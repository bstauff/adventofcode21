package day3

func GetPowerConsumption(diagReport []uint16, resultBitMask uint16) uint64 {
	gamma := getMostCommonBits(diagReport)
	epsilon := uint16(^gamma)
	epsilon = epsilon & resultBitMask

	powerRating := uint64(gamma) * uint64(epsilon)

	return powerRating
}

func GetLifeSupportRating(diagReport []uint16, positionMasks []uint16) uint64 {

	oxygenRating := getOxygenRating(diagReport, positionMasks)

	scrubberRating := getScrubberRating(diagReport, positionMasks)

	liftSupportRating := uint64(oxygenRating) * uint64(scrubberRating)

	return liftSupportRating
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

func getMostCommonBitForPosition(diagReports []uint16, positionMask uint16) uint16 {
	numberOfEntries := len(diagReports)
	countOfSetBits := 0

	for _, report := range diagReports {
		if report&positionMask == positionMask {
			countOfSetBits++
		}
	}

	if countOfSetBits >= numberOfEntries-countOfSetBits {
		return positionMask
	} else {
		return 0
	}
}

func getLeastCommonBitForPosition(diagReports []uint16, positionMask uint16) uint16 {
	numberOfEntries := len(diagReports)
	numberOfOnes := 0

	for _, report := range diagReports {
		if report&positionMask == positionMask {
			numberOfOnes++
		}
	}

	numberOfZeroes := numberOfEntries - numberOfOnes

	if numberOfOnes < numberOfZeroes {
		return positionMask
	} else {
		return 0
	}
}

func getOxygenRating(diagReport []uint16, positionMasks []uint16) uint16 {
	searchSpace := diagReport
	for _, mask := range positionMasks {
		mostCommonBitInPosition := getMostCommonBitForPosition(searchSpace, mask)
		newSearchSpace := make([]uint16, 0, len(searchSpace))
		for _, entry := range searchSpace {
			bitvalue := entry & mask

			shouldGrabReport := bitvalue^(mostCommonBitInPosition&mask) == 0

			if shouldGrabReport {
				newSearchSpace = append(newSearchSpace, entry)
			}
		}
		searchSpace = newSearchSpace

		if len(searchSpace) == 1 {
			break
		}
	}
	return searchSpace[0]
}

func getScrubberRating(diagReport []uint16, positionMasks []uint16) uint16 {
	searchSpace := diagReport
	for _, mask := range positionMasks {
		leastCommonBitInPosition := getLeastCommonBitForPosition(searchSpace, mask)
		newSearchSpace := make([]uint16, 0, len(searchSpace))
		for _, entry := range searchSpace {
			bitvalue := entry & mask

			shouldGrabReport := bitvalue^(leastCommonBitInPosition&mask) == 0

			if shouldGrabReport {
				newSearchSpace = append(newSearchSpace, entry)
			}
		}
		searchSpace = newSearchSpace

		if len(searchSpace) == 1 {
			break
		}
	}
	return searchSpace[0]
}
