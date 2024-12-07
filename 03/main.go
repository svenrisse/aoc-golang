package day3

import (
	"regexp"
	"strconv"
	"strings"
)

type Multiplication struct {
	a, b int
}

func (m Multiplication) Product() int {
	return m.a * m.b
}

func Solve(input string, parseFunc func(string) []Multiplication) int {
	mults := parseFunc(input)
	sum := 0
	for _, m := range mults {
		sum += m.Product()
	}
	return sum
}

func ParseA(input string) []Multiplication {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	result := make([]Multiplication, 0)
	for _, m := range matches {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		result = append(result, Multiplication{a, b})
	}

	return result
}

func ParseB(input string) []Multiplication {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	indexes := r.FindAllStringSubmatchIndex(input, -1)

	result := make([]Multiplication, 0)
	for i, m := range matches {
		start := indexes[i][0]
		if strings.LastIndex(input[:start], "don't()") > strings.LastIndex(input[:start], "do()") {
			continue
		}
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		result = append(result, Multiplication{a, b})

	}
	return result
}
