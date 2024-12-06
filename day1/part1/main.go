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

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "   ")

		i1, _ := strconv.Atoi(split[0])
		left = append(left, i1)

		i2, _ := strconv.Atoi(split[1])
		right = append(right, i2)

	}

	slices.Sort(left)
	slices.Sort(right)

	var total = 0
	for i := 0; i < len(left); i++ {
		total += (Abs(left[i] - right[i]))
	}

	fmt.Println(total)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
