package main

import (
	utils "aoc2024"
	"bufio"
	"log"
	"strings"
)

func a() {
	file := utils.ReadFile("./input.txt")

	scanner := bufio.NewScanner(file)

	countSafeReports := 0

	for scanner.Scan() {
		var splitLine []string
		splitLine = strings.Fields(scanner.Text())

		numbers := utils.StringArrayToInts(splitLine)

		for index, element := range numbers {
			next := numbers[index+1]

			// TODO: checks
			if (element - next <) {
			}
			countSafeReports++
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
