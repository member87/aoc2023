package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var SEARCH_WORD = []string{"X", "M", "A", "S"}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var puzzle [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		puzzle = append(puzzle, strings.Split(scanner.Text(), ""))
	}

	fmt.Println(solve(puzzle))
}

func solve(puzzle [][]string) int {
	var total = 0

	for i := 0; i < len(puzzle); i++ {
		for char := range puzzle[i] {
			if puzzle[i][char] == "X" {
				total += solveWord(puzzle, i, char)
			}
		}
	}

	return total
}

func solveWord(puzzle [][]string, row int, char int) int {
	var total = 0

	total += checkHorizontal(puzzle, row, char)
	total += checkVertical(puzzle, row, char)
	total += checkDiagonal(puzzle, row, char)

	return total
}

func checkHorizontal(puzzle [][]string, row int, char int) int {
	var forward, backward = 1, 1
	for i := 0; i < len(SEARCH_WORD); i++ {
		if forward == 1 && (char+i >= len(puzzle[row]) || puzzle[row][char+i] != SEARCH_WORD[i]) {
			forward = 0
		}
		if backward == 1 && (char-i < 0 || puzzle[row][char-i] != SEARCH_WORD[i]) {
			backward = 0
		}
	}
	return forward + backward
}

func checkVertical(puzzle [][]string, row int, char int) int {
	var down, up = 1, 1
	for i := 0; i < len(SEARCH_WORD); i++ {
		if down == 1 && (row+i >= len(puzzle) || puzzle[row+i][char] != SEARCH_WORD[i]) {
			down = 0
		}
		if up == 1 && (row-i < 0 || puzzle[row-i][char] != SEARCH_WORD[i]) {
			up = 0
		}
	}
	return down + up
}

func checkDiagonal(puzzle [][]string, row int, char int) int {
	var downRight, downLeft, upRight, upLeft = 1, 1, 1, 1
	for i := 0; i < len(SEARCH_WORD); i++ {
		if downRight == 1 && (row+i >= len(puzzle) || char+i >= len(puzzle[row]) || puzzle[row+i][char+i] != SEARCH_WORD[i]) {
			downRight = 0
		}
		if downLeft == 1 && (row+i >= len(puzzle) || char-i < 0 || puzzle[row+i][char-i] != SEARCH_WORD[i]) {
			downLeft = 0
		}
		if upRight == 1 && (row-i < 0 || char+i >= len(puzzle[row]) || puzzle[row-i][char+i] != SEARCH_WORD[i]) {
			upRight = 0
		}
		if upLeft == 1 && (row-i < 0 || char-i < 0 || puzzle[row-i][char-i] != SEARCH_WORD[i]) {
			upLeft = 0
		}
	}

	return downRight + downLeft + upRight + upLeft
}
