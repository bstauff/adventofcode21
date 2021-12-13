package day4

type BingoBoard struct {
	numberToMarkStatus map[int]bool
	numberLocations    map[int][2]int
}

func NewBingoBoard(numbers [][]int) *BingoBoard {
	board := new(BingoBoard)

	board.numberLocations = make(map[int][2]int, 25)
	board.numberToMarkStatus = make(map[int]bool, 25)

	for i := range numbers {
		for j := range numbers[i] {
			board.numberLocations[numbers[i][j]] = [2]int{i, j}
		}
	}
	return board
}

func (board *BingoBoard) MarkNumber(number int) {
	board.numberToMarkStatus[number] = true
}

func (board BingoBoard) HasBoardWon() bool {
	return false
}

func (board BingoBoard) CalculateBoardScore() int {
	return 0
}