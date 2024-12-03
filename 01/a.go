package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
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

	sortAsc(a)
	sortAsc(b)

	fmt.Printf("%v", a)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sortAsc(a []int) {
	sort.Slice(a, func(i int, j int) bool {
		return a[i] < a[j]
	})
}
