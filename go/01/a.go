package main

import (
	utils "aoc2024"
	"bufio"
	"fmt"
	"log"
	"slices"
	"strings"
)

func a() {
	file := utils.ReadFile("./01/input.txt")
	defer file.Close()

	var a []int
	var b []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var splitLine []string
		splitLine = strings.Fields(scanner.Text())

		numbers := utils.StringArrayToInts(splitLine)
		fmt.Println(numbers)

		a = append(a, numbers[0])
		b = append(b, numbers[1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(a)
	slices.Sort(b)

	sum := 0

	for i := 0; i < len(a); i++ {
		difference := a[i] - b[i]
		if difference < 0 {
			difference = -difference
		}
		sum += difference
	}
	println(sum)
}
