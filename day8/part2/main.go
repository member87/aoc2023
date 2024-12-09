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
	valid = cleanValid(valid)
	fmt.Println(len(valid))
}

func cleanValid(valid []string) []string {
	res := []string{}
	for _, v := range valid {
		if !slices.Contains(res, v) {
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
	operations := []string{"+", "-"}

	for _, coord := range others {
		x2, y2 := coord[0], coord[1]
		valid = append(valid, fmt.Sprintf("%d,%d", x1, y1))

		for _, op1 := range operations {
			var px, py int
			if op1 == "+" {
				px = x1 + (x1 - x2)
				py = y1 + (y1 - y2)
			} else {
				px = x1 - (x1 - x2)
				py = y1 - (y1 - y2)
			}

			for true {
				if px >= 0 && px < width && py < height && py >= 0 && !(px == x2 && py == y2) {
					valid = append(valid, fmt.Sprintf("%d,%d", px, py))
					px = px + (x1 - x2)
					py = py + (y1 - y2)
				} else {
					break
				}
			}
		}
	}

	return valid
}
