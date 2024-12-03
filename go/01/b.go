package main

import (
	utils "aoc2024"
	"bufio"
	"log"
	"slices"
	"strconv"
	"strings"
)

func b() {
	file := utils.ReadFile("./input.txt")
	defer file.Close()

	var a []int
	var b []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var splitLine []string
		splitLine = strings.Fields(scanner.Text())

		numA, _ := strconv.Atoi(splitLine[0])
		numB, _ := strconv.Atoi(splitLine[1])

		a = append(a, numA)
		b = append(b, numB)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(a)
	slices.Sort(b)

	similarityScore := 0

	for i := 0; i < len(a); i++ {
		current := a[i]
		counter := 0
	b:
		for j := 0; j < len(b); j++ {
			if b[j] == current {
				counter++
			}
			if b[j] > current {
				similarityScore += counter * current
				counter = 0
				break b
			}
		}
	}

	println(similarityScore)
}
