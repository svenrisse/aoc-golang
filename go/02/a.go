package main

import (
	utils "aoc2024"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func a() {
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

		if isSafe(numbers) {
			countSafeReports++
		}
	}

	fmt.Println(countSafeReports)
}

func isSafee(input []int) bool {
	if input[0] < input[1] {
		for i := 0; i < len(input)-1; i++ {
			if input[i] > input[i+1] || input[i+1]-input[i] > 3 || input[i] == input[i+1] {
				return false
			}
		}
	} else {
		for i := 0; i < len(input)-1; i++ {
			if input[i] < input[i+1] || input[i]-input[i+1] > 3 || input[i] == input[i+1] {
				return false
			}
		}
	}
	return true
}
