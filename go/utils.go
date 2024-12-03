package utils

import (
	"log"
	"os"
	"strconv"
)

func ReadFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func StringArrayToInts(input []string) []int {
	ints := make([]int, len(input))
	for index, a := range input {
		i, err := strconv.Atoi(a)
		if err != nil {
			log.Fatal(err)
		}
		ints[index] = i
	}
	return ints
}
