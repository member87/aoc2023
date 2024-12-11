package main

import (
	"bufio"
	"fmt"
	"os"
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
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		fmt.Println(calc(split, 0))
	}
}

func calc(puzzle []string, count int) int {

	if count == 25 {
		return len(puzzle)
	}

	var res []string

	for _, v := range puzzle {
		if v == "0" {
			res = append(res, "1")
		} else if len(v)%2 == 0 {
			length := len(v)
			left := v[:length/2]
			right := v[length/2:]
			leftInt, _ := strconv.Atoi(left)
			rightInt, _ := strconv.Atoi(right)
			res = append(res, strconv.Itoa(leftInt))
			res = append(res, strconv.Itoa(rightInt))
		} else {
			intVal, _ := strconv.Atoi(v)
			res = append(res, strconv.Itoa(intVal*2024))
		}
	}

	return calc(res, count+1)
}
