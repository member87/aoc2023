package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	var visited []string
	fmt.Println(solve(puzzle, currentRow, currentCol, 0, -1, 0, visited, currentRow, currentCol))
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

func solve(
	puzzle [][]string,
	currentRow int,
	currentCol int,
	x int,
	y int,
	changeCount int,
	visited []string,
	guardStartRow int,
	guardStartCol int,
) int {
	if currentRow+y < 0 || currentRow+y >= len(puzzle) || currentCol+x < 0 || currentCol+x >= len(puzzle[currentRow]) {
		return 0
	}

	total := 0
	coord := fmt.Sprintf("%d,%d,%d,%d", currentCol, currentRow, x, y)
	if changeCount == 0 && puzzle[currentRow+y][currentCol+x] != "#" && !(currentRow == guardStartRow && currentCol == guardStartCol) {
		prev := puzzle[currentRow+y][currentCol+x]
		puzzle[currentRow+y][currentCol+x] = "0"
		total += solve(puzzle, currentRow, currentCol, x, y, 1, visited, guardStartRow, guardStartCol)
		visited = make([]string, 0)
		puzzle[currentRow+y][currentCol+x] = prev
	}

	for true {
		nextChar := puzzle[currentRow+y][currentCol+x]
		if nextChar == "#" || nextChar == "0" {
			if changeCount != 0 {
				changeCount++
			}

			coord = fmt.Sprintf("%d,%d,%d,%d", currentCol, currentRow, x, y)
			if changeCount > 2 && slices.Contains(visited, coord) {
				return 1
			} else if changeCount > 1000 {
				return 0
			}

			if y != 0 {
				x = -y
				y = 0
			} else {
				y = x
				x = 0
			}

			if !slices.Contains(visited, coord) {
				visited = append(visited, coord)
			}

		} else {
			break
		}
	}

	return total + solve(puzzle, currentRow+y, currentCol+x, x, y, changeCount, visited, guardStartRow, guardStartCol)
}

func printPuzzle(puzzle [][]string, currentRow int, currentCol int) {
	prev := puzzle[currentRow][currentCol]
	puzzle[currentRow][currentCol] = "X"

	for i := 0; i < len(puzzle); i++ {
		fmt.Println(strings.Join(puzzle[i], ""))
	}

	puzzle[currentRow][currentCol] = prev
}
