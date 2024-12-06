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

	var safe = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		safe += ProcessLine(scanner.Text())
	}

	fmt.Println(safe)

}

func ProcessLine(line string) int {
	split := strings.Split(line, " ")
	var isIncreasing bool = false
	for i := 0; i < len(split)-1; i++ {
		i1, _ := strconv.Atoi(split[i])
		i2, _ := strconv.Atoi(split[i+1])

		if i == 0 {
			isIncreasing = i1 < i2
		}

		if isIncreasing && i2 <= i1 {
			return 0
		} else if !isIncreasing && i2 >= i1 {
			return 0
		}

		if Abs(i1-i2) > 3 {
			return 0
		}
	}
	return 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
