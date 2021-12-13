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

func (solver BingoSolver) GetScoreOfFirstBoardToWin(numbers []int) (int, error) {
	for _, number := range numbers {
		for _, board := range solver.boards {
			board.MarkNumber(number)
			if board.HasBoardWon() {
				board.CalculateBoardScore(number)
				return board.GetFinalScore(), nil
			}
		}
	}

	return 0, errors.New("no winning board found")
}

func (solver BingoSolver) GetScoreOfLastBoardToWin(numbers []int) (int, error) {
	winningBoards := make(map[int]bool)
	var lastWinningBoard *BingoBoard
	for _, number := range numbers {
		for boardIndex, board := range solver.boards {
			if winningBoards[boardIndex] {
				continue
			}

			board.MarkNumber(number)
			if board.HasBoardWon() {
				board.CalculateBoardScore(number)
				winningBoards[boardIndex] = true
				lastWinningBoard = board
			}
		}
	}

	if lastWinningBoard == nil {
		return 0, errors.New("no winning board found")
	}

	return lastWinningBoard.GetFinalScore(), nil
}
