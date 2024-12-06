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

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var rules []string
	var updates []string
	var savingRules = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			savingRules = false
			continue
		}

		if savingRules {
			rules = append(rules, line)
		} else {
			updates = append(updates, line)
		}
	}

	fmt.Println(getValidUpdates(rules, updates))

}

func getValidUpdates(rules []string, updates []string) int {
	count := 0

	for _, update := range updates {
		count += isValidUpdate(rules, update)
	}

	return count
}

func isValidUpdate(rules []string, update string) int {
	split := strings.Split(update, ",")
	for i := 0; i < len(split); i++ {
		for j := 1 + i; j < len(split); j++ {
			str := split[i] + "|" + split[j]
			if !slices.Contains(rules, str) {
				return 0
			}
		}
	}

	middleIndex := int(math.Floor(float64(len(split) / 2)))

	middle, _ := strconv.Atoi(split[middleIndex])
	return middle
}
