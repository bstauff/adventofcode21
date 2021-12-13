package day4

type BingoBoard struct {
	markedNumbers   []int
	numberLocations map[int][]int
	columnMarks     map[int][]int
	rowMarks        map[int][]int
	finalScore      int
}

func NewBingoBoard(numbers [][]int) *BingoBoard {
	board := new(BingoBoard)

	board.numberLocations = make(map[int][]int)

	board.columnMarks = make(map[int][]int, 5)

	board.rowMarks = make(map[int][]int, 5)

	board.markedNumbers = make([]int, 0, 25)

	for i := range numbers {
		for j := range numbers[i] {
			board.numberLocations[numbers[i][j]] = []int{i, j}
		}
	}
	return board
}

func (board *BingoBoard) MarkNumber(number int) {
	numberLocation := board.numberLocations[number]
	if numberLocation == nil {
		return
	}
	board.markedNumbers = append(board.markedNumbers, number)
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

func (board *BingoBoard) CalculateBoardScore(number int) {
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

	board.finalScore = score * number
}

func (board BingoBoard) GetFinalScore() int {
	return board.finalScore
}
