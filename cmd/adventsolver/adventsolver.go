package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bstauff/adventofcode21/internal/pkg/problemsolvers"
)

func main() {
	inputFromStdIn := make(chan string)
	go startReadingFromStdIn(inputFromStdIn)

	for {
		fmt.Printf("Which day to run?: ")
		selectedDay := <-inputFromStdIn
		if selectedDay == "stop" {
			break
		}
		fmt.Printf("running %s ...\n", selectedDay)
		runDay(selectedDay, inputFromStdIn)
	}
}

func runDay(selectedDay string, inputFromStdIn chan string) {
	switch selectedDay {
	case "Day1SonarSweep":
		runDay1SonarSweep(inputFromStdIn)
	default:
		fmt.Printf("bad day, try again\n")
	}
}

func runDay1SonarSweep(inputFromStdIn chan string) {

	inputDataFomSite, readErr := readDay1InputData()
	if readErr != nil {
		fmt.Printf("failed to read day1data.txt: %s\n", readErr.Error())
		return
	}

	sonarSweeper := new(problemsolvers.Day1SonarSweep)
	initErr := sonarSweeper.Initialize(inputDataFomSite)

	if initErr != nil {
		fmt.Printf("failed to run Day1SonarSweep: %s\n", initErr.Error())
		return
	}

	result, solveErr := sonarSweeper.Solve()

	if solveErr != nil {
		fmt.Printf("failed to solve Day1SonarSweep: %s\n", solveErr.Error())
	}

	fmt.Printf("solution to Day1SonarSweep is: %s\n", result)
}

func readDay1InputData() (string, error) {
	file, err := os.Open("day1data.txt")

	if err != nil {
		return "", err
	}

	defer file.Close()

	var stringBuilder strings.Builder

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		_, writeError := stringBuilder.WriteString(scanner.Text() + ",")
		if writeError != nil {
			return "", writeError
		}
	}
	return strings.TrimSuffix(stringBuilder.String(), ","), nil
}

func startReadingFromStdIn(outputChannel chan string) {
	stdInReader := bufio.NewReader(os.Stdin)
	for {
		valueReadFromStdIn, _ := stdInReader.ReadString('\n')
		outputChannel <- strings.TrimSuffix(valueReadFromStdIn, "\n")
	}
}
