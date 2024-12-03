package utils

import (
	"log"
	"strconv"
)

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
