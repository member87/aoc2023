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
			if puzzle[i][char] == "A" {
				if solveWord(puzzle, i, char) {
					total++
				}
			}
		}
	}

	return total
}

func solveWord(puzzle [][]string, row int, char int) bool {
	var total = 0

	var check = [][][]int{
		{
			{-1, -1},
			{1, 1},
		},
		{
			{1, -1},
			{-1, 1},
		},
	}

	for i := 0; i < len(check); i++ {
		x1, y1 := check[i][0][0], check[i][0][1]
		x2, y2 := check[i][1][0], check[i][1][1]

		if row+y1 < 0 || row+y2 < 0 || row+y1 >= len(puzzle) || row+y2 >= len(puzzle) {
			return false
		}

		if char+x1 < 0 || char+x2 < 0 || char+x1 >= len(puzzle[row]) || char+x2 >= len(puzzle[row]) {
			return false
		}

		char1 := puzzle[row+y1][char+x1]
		char2 := puzzle[row+y2][char+x2]

		if char1 == "M" && char2 == "S" {
			total++
		} else if char1 == "S" && char2 == "M" {
			total++
		}
	}

	return total == 2
}
