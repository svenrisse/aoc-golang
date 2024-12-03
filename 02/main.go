package day2

import "github.com/svenrisse/aoc2024/utils"

func parseInput(input [][]int) []Line {
	var res []Line
	for _, x := range input {
		res = append(res, Line(x))
	}
	return res
}

func solveB(lines []Line) int {
	res := 0
	for _, l := range lines {
		if l.isSafeDamp() {
			res += 1
		}
	}
	return res
}

type Line []int

func (l Line) isSafe() bool {
	return l.isIncreasing() || l.isDecreasing()
}

func (l Line) isIncreasing() bool {
	for index, current := range l {
		if index == 0 {
			continue
		}
		previous := l[index-1]
		if current-previous < 1 || current-previous > 3 {
			return false
		}
	}
	return true
}

func (l Line) isDecreasing() bool {
	for index, current := range l {
		if index == 0 {
			continue
		}
		previous := l[index-1]
		if previous-current < 1 || previous-current > 3 {
			return false
		}
	}
	return true
}

func (l Line) damp() []Line {
	var result []Line
	for i := range l {
		result = append(result, utils.RemoveIndex(l, i))
	}
	return result
}

func (l Line) isSafeDamp() bool {
	for _, dampened := range l.damp() {
		if dampened.isSafe() {
			return true
		}
	}
	return false
}
