package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBingoSolverFindsGoodSolution(t *testing.T) {
	drawnNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	board1Numbers := [][]int{
		[]int{22, 13, 17, 11, 0},
		[]int{8, 2, 23, 4, 24},
		[]int{21, 9, 14, 16, 7},
		[]int{6, 10, 3, 18, 5},
		[]int{1, 12, 20, 15, 19},
	}
	board1 := NewBingoBoard(board1Numbers)

	board2Numbers := [][]int{
		[]int{3, 15, 0, 2, 22},
		[]int{9, 18, 13, 17, 5},
		[]int{19, 8, 7, 25, 23},
		[]int{20, 11, 10, 24, 4},
		[]int{14, 21, 16, 12, 6},
	}
	board2 := NewBingoBoard(board2Numbers)

	board3Numbers := [][]int{
		[]int{14, 21, 17, 24, 4},
		[]int{10, 16, 15, 9, 19},
		[]int{18, 8, 23, 26, 20},
		[]int{22, 11, 13, 6, 5},
		[]int{2, 0, 12, 3, 7},
	}
	board3 := NewBingoBoard(board3Numbers)

	bingoSolver := NewBingoSolver()

	bingoSolver.LoadBoard(board1)
	bingoSolver.LoadBoard(board2)
	bingoSolver.LoadBoard(board3)

	winningScore := bingoSolver.GetWinningBoardScoreFor(drawnNumbers)

	assert.Equal(t, 4512, winningScore)
}
