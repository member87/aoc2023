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

	var left []int
	var right map[int]int = make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "   ")

		i1, _ := strconv.Atoi(split[0])
		left = append(left, i1)

		i2, _ := strconv.Atoi(split[1])
		right[i2] += 1
	}

	var total = 0
	for i := 0; i < len(left); i++ {
		total += (Abs(left[i] * right[left[i]]))
	}

	fmt.Println(total)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
