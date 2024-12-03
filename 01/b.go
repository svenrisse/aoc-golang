package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func b() {
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
