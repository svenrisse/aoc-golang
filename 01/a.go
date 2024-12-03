package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/svenrisse/aoc2024/utils"
)

func a() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

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
