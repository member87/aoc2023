package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var width int

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var puzzle []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		width = len(split)

		for i := 0; i < len(split); i++ {

			intVal, _ := strconv.Atoi(split[i])
			puzzle = append(puzzle, intVal)
		}
	}

	fmt.Println(calc(puzzle))
	areElementsAdjacent(8, 15, 14)
}

func calc(puzzle []int) int {

	total := 0
	for i := 0; i < len(puzzle); i++ {
		if puzzle[i] == 0 {
			res := calculatePath(puzzle, i, 0)
			total += len(getUnique(res))
		}
	}

	return total
}

func getUnique(slice []int) []int {
	var unique []int
	for i := 0; i < len(slice); i++ {
		if !slices.Contains(unique, slice[i]) {
			unique = append(unique, slice[i])
		}
	}
	return unique
}

func areElementsAdjacent(width, i, j int) bool {
	horizontalAdjacent := math.Abs(float64(i-j)) == 1 && (i/width == j/width)
	verticalAdjacent := (i+width == j) || (i-width == j)

	return horizontalAdjacent || verticalAdjacent
}

func calculatePath(puzzle []int, index int, current int) []int {
	operations := []int{1, width, -1, -width}

	if current == 9 {
		return []int{index}
	}

	var res []int

	for i := 0; i < len(operations); i++ {
		nextIndex := index + operations[i]
		if nextIndex < 0 || nextIndex >= len(puzzle) || !areElementsAdjacent(width, index, nextIndex) {
			continue
		}

		if puzzle[nextIndex] == current+1 {
			res = append(res, calculatePath(puzzle, nextIndex, current+1)...)
		}
	}

	return res
}
