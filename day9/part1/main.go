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
		calc(line)
	}
}

func insertIntoSlice(slice []string, value string, count int) []string {
	for i := 0; i < count; i++ {
		slice = append(slice, value)
	}
	return slice
}

func calc(str string) {
	split := strings.Split(str, "")

	var data []string

	for i := 0; i < len(split); i++ {
		intVal, _ := strconv.Atoi(split[i])
		var strToInsert string
		if i%2 == 0 {
			strToInsert = strconv.Itoa(i / 2)
		} else {
			strToInsert = "."
		}
		data = insertIntoSlice(data, strToInsert, intVal)
	}

	data = rearange(data)
	total := checksum(data)

	fmt.Println(total)
}

func rearange(slice []string) []string {
	j := len(slice) - 1

	for i := 0; i < j; i++ {
		if slice[i] == "." {
			slice[i] = slice[j]
			slice[j] = "."

			for true {
				if slice[j] != "." {
					break
				}
				j--
			}
		}
	}

	return slice
}

func checksum(slice []string) int {
	total := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == "." {
			continue
		}
		intVal, _ := strconv.Atoi(slice[i])
		total += intVal * i
	}

	return total
}
