package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardMarkNumberHasWonReturnsWin(t *testing.T) {
	drawnNumbers := []int{8, 2, 23, 4, 24}

	board1Numbers := [][]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}
	board1 := NewBingoBoard(board1Numbers)

	score := 0

	for _, number := range drawnNumbers {
		board1.MarkNumber(number)
		if board1.HasBoardWon() {
			board1.CalculateBoardScore(number)
			score = board1.GetFinalScore()
			break
		}
	}

	assert.Equal(t, 5736, score)
}

func TestBingoSolverFindsGoodSolution(t *testing.T) {
	drawnNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	board1Numbers := [][]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}
	board1 := NewBingoBoard(board1Numbers)

	board2Numbers := [][]int{
		{3, 15, 0, 2, 22},
		{9, 18, 13, 17, 5},
		{19, 8, 7, 25, 23},
		{20, 11, 10, 24, 4},
		{14, 21, 16, 12, 6},
	}
	board2 := NewBingoBoard(board2Numbers)

	board3Numbers := [][]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}
	board3 := NewBingoBoard(board3Numbers)

	bingoSolver := NewBingoSolver()

	bingoSolver.LoadBoard(board1)
	bingoSolver.LoadBoard(board2)
	bingoSolver.LoadBoard(board3)

	winningScore, err := bingoSolver.GetScoreOfFirstBoardToWin(drawnNumbers)
	assert.Nil(t, err)
	assert.Equal(t, 4512, winningScore)
}
func TestBingoSolverFindsGoodLastWinningBoardSolution(t *testing.T) {
	drawnNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	board1Numbers := [][]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}
	board1 := NewBingoBoard(board1Numbers)

	board2Numbers := [][]int{
		{3, 15, 0, 2, 22},
		{9, 18, 13, 17, 5},
		{19, 8, 7, 25, 23},
		{20, 11, 10, 24, 4},
		{14, 21, 16, 12, 6},
	}
	board2 := NewBingoBoard(board2Numbers)

	board3Numbers := [][]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}
	board3 := NewBingoBoard(board3Numbers)

	bingoSolver := NewBingoSolver()

	bingoSolver.LoadBoard(board1)
	bingoSolver.LoadBoard(board2)
	bingoSolver.LoadBoard(board3)

	winningScore, err := bingoSolver.GetScoreOfLastBoardToWin(drawnNumbers)
	assert.Nil(t, err)
	assert.Equal(t, 1924, winningScore)
}
func TestBingoSolverLargeInputFindsGoodSolution(t *testing.T) {

	numbers := []int{42, 32, 13, 22, 91, 2, 88, 85, 53, 87, 37, 33, 76, 98, 89, 19, 69, 9, 62, 21, 38, 49, 54, 81, 0, 26, 79, 36, 57, 18, 4, 40, 31, 80, 24, 64, 77, 97, 70, 6, 73, 23, 20, 47, 45, 51, 74, 25, 95, 96, 58, 92, 94, 11, 39, 63, 65, 99, 48, 83, 29, 34, 44, 75, 55, 17, 14, 56, 8, 82, 59, 52, 46, 90, 5, 41, 60, 67, 16, 1, 15, 61, 71, 66, 72, 30, 28, 3, 43, 27, 78, 10, 86, 7, 50, 35, 84, 12, 93, 68}
	boards, err := ReadInput("input.txt")

	assert.Nil(t, err)
	assert.NotNil(t, boards)

	bingoSolver := NewBingoSolver()

	for _, board := range boards {
		bingoSolver.LoadBoard(board)
	}

	winningScore, err := bingoSolver.GetScoreOfFirstBoardToWin(numbers)

	assert.Nil(t, err)
	assert.NotZero(t, winningScore)
}
func TestBingoSolverFindLastLargeInputFindsGoodSolution(t *testing.T) {

	numbers := []int{42, 32, 13, 22, 91, 2, 88, 85, 53, 87, 37, 33, 76, 98, 89, 19, 69, 9, 62, 21, 38, 49, 54, 81, 0, 26, 79, 36, 57, 18, 4, 40, 31, 80, 24, 64, 77, 97, 70, 6, 73, 23, 20, 47, 45, 51, 74, 25, 95, 96, 58, 92, 94, 11, 39, 63, 65, 99, 48, 83, 29, 34, 44, 75, 55, 17, 14, 56, 8, 82, 59, 52, 46, 90, 5, 41, 60, 67, 16, 1, 15, 61, 71, 66, 72, 30, 28, 3, 43, 27, 78, 10, 86, 7, 50, 35, 84, 12, 93, 68}
	boards, err := ReadInput("input.txt")

	assert.Nil(t, err)
	assert.NotNil(t, boards)

	bingoSolver := NewBingoSolver()

	for _, board := range boards {
		bingoSolver.LoadBoard(board)
	}

	winningScore, err := bingoSolver.GetScoreOfLastBoardToWin(numbers)

	assert.Nil(t, err)
	assert.NotZero(t, winningScore)
}

func ReadInput(filePath string) ([]*BingoBoard, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	boards := make([]*BingoBoard, 0, 100)

	currentBoardNumbers := make([][]int, 5)
	boardNumIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			boards = append(boards, NewBingoBoard(currentBoardNumbers))
			currentBoardNumbers = make([][]int, 5)
			boardNumIndex = 0
			continue
		}
		split := strings.Split(line, " ")
		rowValues := make([]int, 0, 5)
		for _, n := range split {
			if n == "" {
				continue
			}
			number, err := strconv.Atoi(n)

			if err != nil {
				return nil, err
			}
			rowValues = append(rowValues, number)
		}
		currentBoardNumbers[boardNumIndex] = rowValues
		boardNumIndex++
	}

	return boards, nil
}
