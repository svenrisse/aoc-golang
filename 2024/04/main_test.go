package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/svenrisse/aoc2024/utils"
)

var example = Grid{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func TestAllPos(t *testing.T) {
	expected := []Coord{{0, 4}, {0, 5}, {1, 4}}
	assert.Equal(t, expected, example.allPos('X')[:3])
}

func TestRuneAtPos(t *testing.T) {
	rune := example.runeAtPos(Coord{0, 4})
	assert.Equal(t, 'X', rune)

	rune = example.runeAtPos(Coord{0, 6})
	assert.NotEqual(t, 'X', rune)
}

func TestWordFromPos(t *testing.T) {
	val := example.wordFromPos(Coord{0, 4}, Coord{1, 1}, 3)
	assert.Equal(t, "MAS", val)

	val = example.wordFromPos(Coord{0, 4}, Coord{0, -1}, 3)
	assert.NotEqual(t, "MAS", val)

	val = example.wordFromPos(Coord{1, 4}, Coord{0, -1}, 3)
	assert.Equal(t, "MAS", val)

	val = example.wordFromPos(Coord{9, 9}, Coord{0, 1}, 3)
}

func TestCountMAS(t *testing.T) {
	assert.Equal(t, 0, example.countMAS(Coord{0, 3}))
	assert.Equal(t, 1, example.countMAS(Coord{0, 4}))
	assert.Equal(t, 2, example.countMAS(Coord{4, 6}))
	assert.Equal(t, 2, example.countMAS(Coord{9, 9}))
}

func TestSolveAExample(t *testing.T) {
	assert.Equal(t, 18, solveA(example))
}

func TestSolveA(t *testing.T) {
	grid := Grid(utils.LoadStrings("input.txt"))
	result := solveA(grid)
	assert.Equal(t, 2517, result)
}

func TestSolveBExample(t *testing.T) {
	assert.Equal(t, 9, solveB(example))
}

func TestSolveB(t *testing.T) {
	grid := Grid(utils.LoadStrings("input.txt"))
	result := solveB(grid)
	assert.Equal(t, 1960, result)
}
