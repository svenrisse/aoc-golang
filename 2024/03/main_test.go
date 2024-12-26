package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/svenrisse/aoc2024/utils"
)

var exampleA = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func TestParseA(t *testing.T) {
	res := ParseA(exampleA)
	expected := []Multiplication{
		{2, 4}, {5, 5}, {11, 8}, {8, 5},
	}
	assert.Equal(t, expected, res)
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 161, Solve(exampleA, ParseA))
}

func TestSolveA(t *testing.T) {
	input := utils.LoadString("input.txt")
	assert.Equal(t, 182780583, Solve(input, ParseA))
}

var exampleB = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func TestParseB(t *testing.T) {
	res := ParseB(exampleB)
	assert.Equal(t, []Multiplication{{2, 4}, {5, 5}, {11, 8}, {8, 5}}, res)
}

func TestSolveB(t *testing.T) {
	input := utils.LoadString("input.txt")
	assert.Equal(t, 90772405, Solve(input, ParseB))
}
