package main

import (
	utils "aoc2024"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	file := utils.ReadFile("./02/input.txt")

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	countSafeReports := 0

	for scanner.Scan() {
		var splitLine []string
		splitLine = strings.Fields(scanner.Text())

		numbers := utils.StringArrayToInts(splitLine)

		if isSafeWithDampener(numbers, false) {
			countSafeReports++
		}
	}

	fmt.Println(countSafeReports)
}

func isSafeWithDampener(input []int, dampenend bool) bool {
	if input[0] < input[1] {
		for i := 0; i < len(input)-1; i++ {
			if input[i] > input[i+1] || input[i+1]-input[i] > 3 || input[i] == input[i+1] {
				if dampenend {
					return false
				}
				return isSafeWithDampener(utils.RemoveElement(input, i+1), true)
			}
		}
	} else {
		for i := 0; i < len(input)-1; i++ {
			if input[i] < input[i+1] || input[i]-input[i+1] > 3 || input[i] == input[i+1] {
				if dampenend {
					return false
				}
				return isSafeWithDampener(utils.RemoveElement(input, i+1), true)
			}
		}
	}
	return true
}
