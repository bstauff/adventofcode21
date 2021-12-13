package day4

type BingoBoard struct {
	numbers [][]int
}

func NewBingoBoard(numbers [][]int) *BingoBoard {
	board := new(BingoBoard)
	board.numbers = numbers
	return board
}