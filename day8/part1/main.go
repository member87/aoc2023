package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	locations := make(map[string][][]int)
	puzzle := [][]string{}

	row := 0
	col := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		if col == 0 {
			col = len(split)
		}
		puzzle = append(puzzle, split)

		for k, v := range split {
			if v != "." {
				value, exists := locations[v]
				coord := []int{k, row}

				if !exists {
					value = [][]int{}
				}

				locations[v] = append(value, coord)
			}
		}
		row++
	}

	valid := slices.Compact(calc(locations, col, row))
	valid = cleanValid(puzzle, valid)
	fmt.Println(len(valid))
}

func cleanValid(puzzle [][]string, valid []string) []string {
	res := []string{}
	for _, v := range valid {
		split := strings.Split(v, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		if puzzle[y][x] == "." {
			res = append(res, v)
		}
	}
	return res
}

func calc(locations map[string][][]int, width int, height int) []string {
	valid := []string{}

	for _, v := range locations {
		for _, coord := range v {
			var others [][]int
			for _, other := range v {
				if other[0] != coord[0] && other[1] != coord[1] {
					others = append(others, other)
				}
			}

			valid = append(valid, findPoints(coord, others, width, height)...)
		}
	}

	return valid
}

func findPoints(initalCoord []int, others [][]int, width int, height int) []string {

	x1, y1 := initalCoord[0], initalCoord[1]

	valid := []string{}

	for _, coord := range others {
		x2, y2 := coord[0], coord[1]
		px1, py1 := x1+(x1-x2), y1+(y1-y2)
		px2, py2 := x1+(x1-x2), y1+(y1-y2)

		if px1 >= 0 && px1 < width && py1 < height && py1 >= 0 && !(px1 == x2 && py1 == y2) {
			valid = append(valid, fmt.Sprintf("%d,%d", px1, py1))
		}

		if px2 >= 0 && px2 < width && py2 < height && py2 >= 0 && !(px2 == x2 && py2 == y2) {
			valid = append(valid, fmt.Sprintf("%d,%d", px2, py2))
		}
	}

	return valid
}
