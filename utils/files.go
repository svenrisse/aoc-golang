package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadString(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func LoadStrings(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func LoadInts(filename string) []int {
	var result []int
	for _, line := range LoadStrings(filename) {
		lineInt, _ := strconv.Atoi(line)
		result = append(result, lineInt)
	}
	return result
}

func LoadIntCols(filename string) [][]int {
	var result [][]int
	for _, line := range LoadStrings(filename) {
		fields := strings.Fields(line)
		var lineInts []int
		for _, f := range fields {
			i, _ := strconv.Atoi(f)
			lineInts = append(lineInts, i)
		}
		result = append(result, lineInts)
	}
	return result
}
