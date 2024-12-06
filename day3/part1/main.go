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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += ProcessLine(scanner.Text())
	}

	fmt.Println(total)
}

func ProcessLine(line string) int {

	r, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")

	var total = 0

	for _, match := range r.FindAllString(line, -1) {
		split := strings.Split(match[4:len(match)-1], ",")

		i1, _ := strconv.Atoi(split[0])
		i2, _ := strconv.Atoi(split[1])

		total += i1 * i2
	}

	return total
}
