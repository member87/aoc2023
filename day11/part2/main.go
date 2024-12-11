package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache map[string]int

func main() {
	file, err := os.Open("input.txt")

	cache = make(map[string]int)

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

	if count == 75 {
		return 1
	}

	total := 0
	for _, v := range puzzle {
		cacheKey := fmt.Sprintf("%s-%d", v, count)
		if v == "0" {
			cache, ok := getCache(cacheKey)
			if ok {
				total += cache
			} else {
				total += setCache(cacheKey, calc([]string{"1"}, count+1))
			}
		} else if len(v)%2 == 0 {
			length := len(v)
			left := v[:length/2]
			right := v[length/2:]
			cache, ok := getCache(cacheKey)
			if ok {
				total += cache
			} else {
				leftInt, _ := strconv.Atoi(left)
				rightInt, _ := strconv.Atoi(right)
				res := calc([]string{strconv.Itoa(leftInt)}, count+1) + calc([]string{strconv.Itoa(rightInt)}, count+1)
				total += setCache(cacheKey, res)
			}
		} else {
			cache, ok := getCache(cacheKey)
			if ok {
				total += cache
			} else {
				intVal, _ := strconv.Atoi(v)
				total += calc([]string{strconv.Itoa(intVal * 2024)}, count+1)
			}
		}
	}

	return total
}

func setCache(key string, val int) int {
	cache[key] = val
	return val
}

func getCache(key string) (int, bool) {
	val, ok := cache[key]
	if ok {
		return val, true
	}
	return 0, false
}
