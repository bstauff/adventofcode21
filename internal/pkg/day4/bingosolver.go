package day4

import "errors"

type BingoSolver struct {
	boards []*BingoBoard
}

func NewBingoSolver() *BingoSolver {
	solver := new(BingoSolver)
	return solver
}

func (solver *BingoSolver) LoadBoard(board *BingoBoard) {
	solver.boards = append(solver.boards, board)
}

func (solver BingoSolver) GetWinningBoardScoreFor(numbers []int) (int, error) {
	for _, number := range numbers {
		for _, board := range solver.boards {
			board.MarkNumber(number)
			if board.HasBoardWon() {
				return board.CalculateBoardScore(), nil
			}
		}
	}

	return 0, errors.New("no winning board found")
}
