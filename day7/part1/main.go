package main

import (
	"bufio"
	"fmt"
	"math"
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
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		total += solve(split...)
	}

	fmt.Println(total)
}

func convertToIntSlice(s []string) []int {
	var result []int
	for _, v := range s {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}
	return result
}

func getBitVal(number int, position int) int {
	if (number & (1 << position)) > 0 {
		return 1
	}
	return 0
}

func solve(s ...string) int {

	answerStr := s[0][:len(s[0])-1]

	answer, _ := strconv.Atoi(answerStr)
	others := convertToIntSlice(s[1:])

	for i := 0; i < int(math.Pow(2, float64(len(others)-1))); i++ {
		var operations []string
		for j := 0; j < len(others)-1; j++ {
			if getBitVal(i, j) == 1 {
				operations = append(operations, "+")
			} else {
				operations = append(operations, "*")
			}
		}

		res := calc(others, operations)
		if res == answer {
			return res
		}
	}

	return 0
}

func calc(numbers []int, operations []string) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if operations[i-1] == "+" {
			result += numbers[i]
		} else {
			result *= numbers[i]
		}
	}
	return result
}
