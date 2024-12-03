package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func a() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

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
