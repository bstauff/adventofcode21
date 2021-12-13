package day4

type BingoBoard struct {
	markedNumbers []int
	numberLocations    map[int][2]int
	columnMarks map[int][]int
	rowMarks map[int][]int
}

func NewBingoBoard(numbers [][]int) *BingoBoard {
	board := new(BingoBoard)

	board.numberLocations = make(map[int][2]int, 25)

	board.columnMarks = make(map[int][]int, 5)
	for i := range board.columnMarks {
		board.columnMarks[i] = make([]int, 5)
	}

	board.rowMarks = make(map[int][]int, 5)
	for i := range board.rowMarks {
		board.rowMarks[i] = make([]int, 5)
	}

	board.markedNumbers = make([]int,0, 25)

	for i := range numbers {
		for j := range numbers[i] {
			board.numberLocations[numbers[i][j]] = [2]int{i, j}
		}
	}
	return board
}

func (board *BingoBoard) MarkNumber(number int) {
	board.markedNumbers = append(board.markedNumbers, number)
	numberLocation := board.numberLocations[number]
	board.columnMarks[numberLocation[1]] = append(board.columnMarks[numberLocation[1]], number)
	board.rowMarks[numberLocation[0]] = append(board.rowMarks[numberLocation[0]], number)
}

func (board BingoBoard) HasBoardWon() bool {
	for _, marks := range board.columnMarks {
		if len(marks) == 5 {
			return true
		}
	}
	for _, marks := range board.rowMarks {
		if len(marks) == 5 {
			return true
		}
	}
	return false
}

func (board BingoBoard) CalculateBoardScore() int {
	isMarked := make(map[int]bool, len(board.markedNumbers))
	for _, number := range board.markedNumbers {
		isMarked[number] = true
	}

	score := 0

	for k := range board.numberLocations {
		if !isMarked[k] {
			score += k
		}
	}

	return score
}