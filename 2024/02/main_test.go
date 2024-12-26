package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/svenrisse/aoc2024/utils"
)

var example = []Line{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 4, solveB(example))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadIntCols("input.txt")
	assert.Equal(t, 271, solveB(parseInput(input)))
}
