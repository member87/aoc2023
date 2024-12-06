package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var puzzle [][]string
	var currentRow int
	var currentCol int

	lineNumber := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		lineString := strings.Split(line, "")
		puzzle = append(puzzle, lineString)

		if currentRow == 0 && strings.Contains(line, "^") {
			currentRow = lineNumber
			currentCol = strings.Index(line, "^")
		}
		lineNumber++
	}

	fmt.Println(solve(puzzle, currentRow, currentCol, 0, -1))
}

func countChar(s []string, c string) int {
	count := 0
	for _, v := range s {
		if v == c {
			count++
		}
	}
	return count
}

func solve(puzzle [][]string, currentRow int, currentCol int, x int, y int) int {
	if currentRow+y < 0 || currentRow+y >= len(puzzle) || currentCol+x < 0 || currentCol+x >= len(puzzle[currentRow]) {
		puzzle[currentRow][currentCol] = "X"
		return 1
	}

	nextChar := puzzle[currentRow+y][currentCol+x]
	currentChar := puzzle[currentRow][currentCol]

	if nextChar == "#" {
		if y != 0 {
			x = -y
			y = 0
		} else {
			y = x
			x = 0
		}
	}

	if currentChar != "X" {
		puzzle[currentRow][currentCol] = "X"
		return 1 + solve(puzzle, currentRow+y, currentCol+x, x, y)
	}

	return solve(puzzle, currentRow+y, currentCol+x, x, y)
}
