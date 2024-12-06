package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var total = 0
	var enabled = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total_res, enabled_res := ProcessLine(scanner.Text(), enabled)
		total += total_res
		enabled = enabled_res
	}

	fmt.Println(total)
}

func ProcessLine(line string, enabled bool) (int, bool) {

	r, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)|do\\(\\)|don't\\(\\)")

	var total = 0

	for _, match := range r.FindAllString(line, -1) {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else if enabled {
			split := strings.Split(match[4:len(match)-1], ",")

			i1, _ := strconv.Atoi(split[0])
			i2, _ := strconv.Atoi(split[1])

			total += i1 * i2
		}
	}

	return total, enabled
}
