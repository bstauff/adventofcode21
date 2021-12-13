package day4

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

func (solver BingoSolver) GetWinningBoardScoreFor(numbers []int) int {
	return 0
}